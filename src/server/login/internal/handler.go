package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {

}

func getLoginByName(username string) (bool, error) {
	
}
func handleUserRegister(args []interface{}) {
	m := args[0].(*msg.RegisterRequest)
	a := args[1].(gate.Agent)
	username := m.GetAccount()
	password := m.GetPassword()
	log.Debug("receive UserRegister LoginName: " + username)
	ok, err := getLoginByName(username)
	if err != nil {
		log.Println(err)
	}
	if ok == true {
		logging.LoginLogger.Debug("UserRegister find in fail")

		retBuf := &msg.DefaultReault{
			RetResult: e.ACCOUNT_EXIST,
			ErrorInfo: e.GetMsg(e.ACCOUNT_EXIST),
		}

		a.WriteMsg(retBuf)
	}
	err = models.AddLogin(name, pwd)
	if err != nil {
		log.Println(err)
		logging.LoginLogger.Debug("UserRegister write in fail")
		retBuf := &msg.DefaultReault{
			RetResult: e.REGISTRE_FAIL,
			ErrorInfo: e.GetMsg(e.REGISTRE_FAIL),
		}
		a.WriteMsg(retBuf)
	} else {
		logging.LoginLogger.Info("UserRegister write in success")
		retBuf := &msg.DefaultReault{
			RetResult: e.REGISTRE_SUCCESS,
		}

		a.WriteMsg(retBuf)
	}
}
