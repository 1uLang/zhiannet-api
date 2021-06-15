package request

//登陆获取cookie
//func Login(req *ApiKey) (string, error) {
//	var err error
//	// https://182.131.30.171:28443/cgi-bin/login.cgi
//	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
//	resp, err := client.R().
//		SetHeader("Content-Type", "multipart/form-data; boundary=<calculated when request is sent>").
//		SetFormData(map[string]string{
//			"param_type":     "login",
//			"param_username": req.Name,
//			"param_password": req.Password,
//		}).
//		Post("https://" + req.Addr + ":" + req.Port + _const.DDOS_LOGIN_URL)
//	if err != nil {
//		logrus.Error(err)
//		return Cookie, err
//	}
//
//	//logrus.Info(string(resp.Body()))
//	var data = LoginRes{}
//	err = xml.Unmarshal(resp.Body(), &data)
//	if err != nil {
//		logrus.Error(err)
//		return Cookie, err
//	}
//	//logrus.Info(data)
//	if data.Failure.Info != "" {
//		logrus.Debug(data.Failure)
//		err = fmt.Errorf(data.Failure.Info)
//		return Cookie, err
//	}
//	if len(resp.Cookies()) > 0 {
//		cook := resp.Cookies()[0]
//		Cookie = cook.Value
//	}
//	fmt.Println(Cookie)
//
//	return Cookie, err
//	//logrus.Info( err)
//}
//
//func GetCookie(req *LoginReq) (cookie string) {
//
//	key := fmt.Sprintf("ddos-cookie-%v:%v", req.Addr, req.Port)
//	//cache.CheckCache(key, Login(req), 3600, true)
//	res, err := cache.GetCache( key)
//	if err != nil {
//		if err == redis.Nil {
//			cookie, _ = Login(req)
//			cache.SetCache( key, cookie, 3600)
//		}
//		return
//	}
//	cookie = fmt.Sprintf("%v", res)
//	return
//}
