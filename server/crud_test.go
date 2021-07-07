package server

// func TestCreateCrypto(t *testing.T) {
// 	connectionDb()
// 	var conn *grpc.ClientConn

// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Could not connect: %s", err)
// 	}

// 	defer conn.Close()

// 	c := system.NewUpVoteServiceClient(conn)

// 	defer clearDB()

// 	crypto := system.Cryptocurrency{
// 		Name:        "",
// 		Initials:    "",
// 		Description: "",
// 	}

// 	//create
// 	createCrypto := system.CreateCryptocurrencyRequest{Crypto: &crypto}

// 	response, err := c.CreateCryptocurrency(context.Background(), &createCrypto)

// 	require.NotNil(t, err)
// 	require.Nil(t, response)

// 	assert.Equal(t, "rpc error: code = InvalidArgument desc = Empty fields", err.Error())

// 	crypto = system.Cryptocurrency{
// 		Name:        "Bitcoin",
// 		Initials:    "BTC",
// 		Description: "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)",
// 	}

// 	// Should to be able valid request
// 	validRequest := system.CreateCryptocurrencyRequest{Crypto: &crypto}

// 	response, err = c.CreateCryptocurrency(context.Background(), &validRequest)
// 	require.Nil(t, err)

// 	assert.Equal(t, "Bitcoin", response.GetCrypto().GetName())
// 	assert.Equal(t, "BTC", response.GetCrypto().GetInitials())
// 	assert.Equal(t, "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)", response.GetCrypto().GetDescription())
// 	assert.Equal(t, int32(0), response.GetCrypto().GetUpvote())
// 	assert.Equal(t, int32(0), response.GetCrypto().GetDownvote())

// 	// SHould to be able valid but the crypto already exists
// 	response, err = c.CreateCryptocurrency(context.Background(), &validRequest)

// 	require.NotNil(t, err)
// 	require.Nil(t, response)
// 	assert.Equal(t, "rpc error: code = AlreadyExists desc = Cryptocurrency already exists", err.Error())

// }

// func TestReadCryptoByID(t *testing.T) {
// 	connectionDb()
// 	var conn *grpc.ClientConn

// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Could not connect: %s", err)
// 	}

// 	defer conn.Close()

// 	c := system.NewUpVoteServiceClient(conn)

// 	defer clearDB()

// 	// Test request with empty ID
// 	emptyIDRequest := &system.ReadCryptocurrencyRequest{
// 		Id: "",
// 	}
// 	response, err := c.ReadCryptocurrencyById(context.Background(), emptyIDRequest)

// 	require.NotNil(t, err)
// 	require.Nil(t, response)

// 	assert.Equal(t, "rpc error: code = InvalidArgument desc = the provided hex string is not a valid ObjectID", err.Error())

// 	// Test request with valid ID but not found on DB
// 	NotFoundIDRequest := &system.ReadCryptocurrencyRequest{
// 		Id: primitive.NewObjectID().Hex(),
// 	}

// 	response, err = c.ReadCryptocurrencyById(context.Background(), NotFoundIDRequest)

// 	require.Nil(t, err)
// 	require.NotNil(t, response)

// 	// Test with valid request
// 	createRequest := system.CreateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Name:        "Bitcoin",
// 			Initials:    "BTC",
// 			Description: "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)",
// 		},
// 	}

// 	cryptoResponse, err := c.CreateCryptocurrency(context.Background(), &createRequest)

// 	require.Nil(t, err)

// 	validRequest := system.ReadCryptocurrencyRequest{
// 		Id: cryptoResponse.Crypto.GetId(),
// 	}

// 	response, err = c.ReadCryptocurrencyById(context.Background(), &validRequest)

// 	require.NotNil(t, response)
// 	require.Nil(t, err)

// 	assert.Equal(t, cryptoResponse.GetCrypto().GetId(), response.GetCrypto().GetId())
// 	assert.Equal(t, cryptoResponse.GetCrypto().GetName(), response.GetCrypto().GetName())
// 	assert.Equal(t, cryptoResponse.GetCrypto().GetInitials(), response.GetCrypto().GetInitials())
// 	assert.Equal(t, cryptoResponse.GetCrypto().GetDescription(), response.GetCrypto().GetDescription())
// 	assert.Equal(t, cryptoResponse.GetCrypto().GetUpvote(), response.GetCrypto().GetUpvote())
// 	assert.Equal(t, cryptoResponse.GetCrypto().GetDownvote(), response.GetCrypto().GetDownvote())

// }

// func TestReadAllCrypto(t *testing.T) {
// 	connectionDb()
// 	var conn *grpc.ClientConn

// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Could not connect: %s", err)
// 	}

// 	defer conn.Close()

// 	c := system.NewUpVoteServiceClient(conn)

// 	defer clearDB()

// 	// Test with no crypto found on DB
// 	var allCreatedCrypto []*system.CreateCryptocurrencyResponse

// 	client := system.NewUpVoteServiceClient(conn)
// 	request := &system.ListAllCryptocurrenciesRequest{}

// 	stream, err := client.ListAllCriptocurrencies(context.Background(), request)

// 	var result []*system.ListAllCryptocurrenciesResponse
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			resp, err := stream.Recv()
// 			if err == io.EOF {
// 				done <- true
// 				return
// 			}
// 			require.Nil(t, err)

// 			result = append(result, resp)

// 		}
// 	}()

// 	<-done

// 	require.Nil(t, err)
// 	require.Nil(t, result)

// 	a := system.Cryptocurrency{
// 		Name:        "Bitcoin",
// 		Initials:    "BTC",
// 		Description: "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)",
// 	}

// 	b := system.Cryptocurrency{
// 		Name:        "Ethereum",
// 		Initials:    "ETH",
// 		Description: "Ethereum is a decentralized open-source blockchain system that features its own cryptocurrency (...)",
// 	}

// 	d := system.Cryptocurrency{
// 		Name:        "Tether",
// 		Initials:    "USDT",
// 		Description: "USDT is a stablecoin (stable-value cryptocurrency) that mirrors the price of the U.S. dollar (...)",
// 	}

// 	// Test with populated DB
// 	cryptoTest := []*system.Cryptocurrency{&a, &b, &d}

// 	for _, crypto := range cryptoTest {
// 		createRequest := &system.CreateCryptocurrencyRequest{Crypto: crypto}
// 		response, err := c.CreateCryptocurrency(context.Background(), createRequest)
// 		require.Nil(t, err)
// 		allCreatedCrypto = append(allCreatedCrypto, response)
// 		require.NotNil(t, allCreatedCrypto)
// 	}

// 	stream, err = client.ListAllCriptocurrencies(context.Background(), request)

// 	for {
// 		resp, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		require.Nil(t, err)

// 		result = append(result, resp)

// 	}

// 	require.Nil(t, err)
// 	for i, crypto := range cryptoTest {
// 		assert.Equal(t, crypto.GetName(), result[i].GetCrypto().GetName())
// 		assert.Equal(t, crypto.GetInitials(), result[i].GetCrypto().GetInitials())
// 		assert.Equal(t, crypto.GetDescription(), result[i].GetCrypto().GetDescription())
// 		assert.Equal(t, crypto.GetUpvote(), result[i].GetCrypto().GetUpvote())
// 		assert.Equal(t, crypto.GetDownvote(), result[i].GetCrypto().GetDownvote())
// 	}

// }

// func TestDeleteCrypto(t *testing.T) {
// 	connectionDb()
// 	var conn *grpc.ClientConn

// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Could not connect: %s", err)
// 	}

// 	defer conn.Close()

// 	c := system.NewUpVoteServiceClient(conn)

// 	defer clearDB()

// 	// Test request with empty ID
// 	emptyIDRequest := &system.DeleteCryptocurrencyRequest{
// 		Id: "",
// 	}
// 	_, err = c.DeleteCryptocurrency(context.Background(), emptyIDRequest)

// 	require.NotNil(t, err)

// 	assert.Equal(t, "rpc error: code = InvalidArgument desc = the provided hex string is not a valid ObjectID", err.Error())

// 	// Test request with valid ID but not found on DB
// 	NotFoundIDRequest := system.DeleteCryptocurrencyRequest{
// 		Id: primitive.NewObjectID().Hex(),
// 	}

// 	_, err = c.DeleteCryptocurrency(context.Background(), &NotFoundIDRequest)

// 	require.Nil(t, err)

// 	// Test with valid request
// 	createRequest := &system.CreateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Name:        "Bitcoin",
// 			Initials:    "BTC",
// 			Description: "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)",
// 		},
// 	}

// 	cryptoResponse, err := c.CreateCryptocurrency(context.Background(), createRequest)

// 	require.Nil(t, err)

// 	validRequest := &system.DeleteCryptocurrencyRequest{
// 		Id: cryptoResponse.Crypto.GetId(),
// 	}

// 	response, err := c.DeleteCryptocurrency(context.Background(), validRequest)

// 	require.Nil(t, err)

// 	assert.Equal(t, true, response.GetStatus())

// }

// func TestUpdateCrypto(t *testing.T) {
// 	connectionDb()
// 	var conn *grpc.ClientConn

// 	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Could not connect: %s", err)
// 	}

// 	defer conn.Close()

// 	c := system.NewUpVoteServiceClient(conn)

// 	defer clearDB()

// 	// Test request with empty ID
// 	emptyIDRequest := &system.UpdateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Id: "",
// 		},
// 	}
// 	_, err = c.UpdateCryptocurrency(context.Background(), emptyIDRequest)

// 	require.NotNil(t, err)

// 	assert.Equal(t, "rpc error: code = InvalidArgument desc = the provided hex string is not a valid ObjectID", err.Error())

// 	// Test request with empty fields
// 	emptyFieldsRequest := &system.UpdateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Id:          primitive.NewObjectID().Hex(),
// 			Name:        "",
// 			Description: "",
// 		},
// 	}
// 	_, err = c.UpdateCryptocurrency(context.Background(), emptyFieldsRequest)

// 	require.NotNil(t, err)

// 	assert.Equal(t, "rpc error: code = InvalidArgument desc = Empty fields", err.Error())

// 	// Test with valid request but ID not found on DB
// 	NotFoundIDRequest := &system.UpdateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Id:          primitive.NewObjectID().Hex(),
// 			Name:        "Bitcoin",
// 			Initials:    "BTC",
// 			Description: "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)",
// 		},
// 	}
// 	_, err = c.UpdateCryptocurrency(context.Background(), NotFoundIDRequest)

// 	require.Nil(t, err)

// 	// Test with valid request
// 	createRequest := &system.CreateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Name:        "Bitcoin",
// 			Initials:    "BTC",
// 			Description: "Bitcoin is a decentralized cryptocurrency originally described in a 2008 (...)",
// 		},
// 	}

// 	cryptoResponse, err := c.CreateCryptocurrency(context.Background(), createRequest)

// 	require.Nil(t, err)

// 	validRequest := system.UpdateCryptocurrencyRequest{
// 		Crypto: &system.Cryptocurrency{
// 			Id:          cryptoResponse.GetCrypto().GetId(),
// 			Name:        "Ethereum",
// 			Initials:    "BTC",
// 			Description: "Ethereum is a decentralized open-source blockchain system that features its own cryptocurrency (...)",
// 		},
// 	}

// 	response, err := c.UpdateCryptocurrency(context.Background(), &validRequest)

// 	require.Nil(t, err)

// 	assert.Equal(t, validRequest.GetCrypto().GetId(), response.GetCrypto().GetId())
// 	assert.Equal(t, validRequest.GetCrypto().GetInitials(), response.GetCrypto().GetInitials())
// 	assert.Equal(t, validRequest.GetCrypto().GetName(), response.GetCrypto().GetName())
// 	assert.Equal(t, validRequest.GetCrypto().GetDescription(), response.GetCrypto().GetDescription())
// 	assert.Equal(t, validRequest.GetCrypto().GetUpvote(), response.GetCrypto().GetUpvote())
// 	assert.Equal(t, validRequest.GetCrypto().GetDownvote(), response.GetCrypto().GetDownvote())

// }
