package controllers

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"encoding/binary"

	"github.com/gin-gonic/gin"
	"github.com/brianosaurus/volume-practice/api/utils/formaterror"
)


//CreateEstate makes a estate
func (server *Server) FlightPaths(c *gin.Context) {
	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	estate := models.Estate{}

	err = json.Unmarshal(body, &estate)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	estate.Prepare()

	errorMessages := estate.Validate("")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}
	estate.Owned = false
	estateCreated, err := estate.SaveEstate(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	// create the signature needed to verify we have created the least for the estate
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	publicKey := privateKey.Public()

	fmt.Println("YYYYYYYYYYYY")
	fmt.Println(publicKey)
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) 
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	fmt.Println("ZZZZZZZZZZZZZZZZ")


	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		errList = formattedError
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
    }

	fmt.Println(estate.Owner)

	topLeftX := make([]byte, 4)
	topLeftY := make([]byte, 4)
	bottomRightX := make([]byte, 4)
	bottomRightY := make([]byte, 4)

	binary.BigEndian.PutUint32(topLeftX, estate.TopLeftX)
	binary.BigEndian.PutUint32(topLeftY, estate.TopLeftY)
	binary.BigEndian.PutUint32(bottomRightX, estate.BottomRightX)
	binary.BigEndian.PutUint32(bottomRightY, estate.BottomRightY)

	// pack and hash data
	hash := crypto.Keccak256Hash(
        common.HexToAddress(estate.Owner).Bytes(),
        common.LeftPadBytes(estate.BlockNo.Bytes(), 32),
        topLeftX,
        topLeftY,
        bottomRightX,
        bottomRightY,
    )
    // normally we sign prefixed hash
    // as in solidity with `ECDSA.toEthSignedMessageHash`
	hash = crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n32")),
		hash.Bytes(),
	)

	// sign hash to validate later in Solidity
	// sig, err := crypto.Sign(prefixedHash.Bytes(), privateKey)
	// fmt.Println(privateKey)
	sig, err := crypto.Sign(hash.Bytes(), privateKey)

	// this is to make the signature the 'legacy' style which the ethereum blockchain uses
	// figuring that out wasn't easy
	sig[crypto.RecoveryIDOffset] += 27

	c.JSON(http.StatusCreated, gin.H{
		"signature": hexutil.Encode(sig),
		"blockNo": estate.BlockNo,
		"status":   http.StatusCreated,
		"response": estateCreated,
	})
}

//GetEstates lists all estates
func (server *Server) GetEstates(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	estate := models.Estate{}

	estates, err := estate.FindAllEstates(server.DB)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": estates,
	})
}

//GetEstate gets a specific estate
func (server *Server) GetEstate(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	estateID := c.Param("id")

	hid, err := strconv.ParseUint(estateID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	estate := models.Estate{}

	estateGotten, err := estate.FindEstateByID(server.DB, hid)
	if err != nil {
		errList["No_user"] = "No User Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": estateGotten,
	})
}

//UpdateEstate updates a single estate
func (server *Server) UpdateEstate(c *gin.Context) {
	//clear previous error if any
	errList = map[string]string{}

	estateID := c.Param("id")
	// Check if the user id is valid
	hid, err := strconv.ParseUint(estateID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}
	
	// Start processing the request
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errList["Invalid_body"] = "Unable to get request"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	estate := &models.Estate{}
	err = json.Unmarshal(body, &estate)
	if err != nil {
		errList["Unmarshal_error"] = "Cannot unmarshal body"
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	estate.Prepare()
	errorMessages := estate.Validate("update")
	if len(errorMessages) > 0 {
		errList = errorMessages
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errList,
		})
		return
	}

	estate.ID = hid

	updatedEstate, err := estate.UpdateEstate(server.DB, hid)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": updatedEstate,
	})
}

//DeleteEstate removes a estate from the DB
func (server *Server) DeleteEstate(c *gin.Context) {
	estateID := c.Param("id")
	// Is a valid post id given to us?
	hid, err := strconv.ParseUint(estateID, 10, 64)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	// Check if the comment exist
	estate := models.Estate{}
	err = server.DB.Debug().Model(models.Estate{}).Where("id = ?", hid).Take(&estate).Error
	if err != nil {
		errList["No_post"] = "No Post Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}

	// If all the conditions are met, delete the post
	_, err = estate.DeleteEstate(server.DB)
	if err != nil {
		errList["Other_error"] = "Please try again later"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "Comment deleted",
	})
}
