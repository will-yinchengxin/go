package mapStruct

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
)

type UserType struct {
	UserTypeId   int
	UserTypeName string
}

type NumberFormat struct {
	DecimalSeparator  string `jpath:"userContext.conversationCredentials.sessionToken"`
	GroupingSeparator string `jpath:"userContext.preferenceInfo.numberFormat.groupingSeparator"`
	GroupPattern      string `jpath:"userContext.preferenceInfo.numberFormat.groupPattern"`
}

type User struct {
	Session      string   `jpath:"userContext.cobrandConversationCredentials.sessionToken"`
	CobrandId    int      `jpath:"userContext.cobrandId"`
	UserType     UserType `jpath:"userType"`
	LoginName    string   `jpath:"loginName"`
	NumberFormat          // This can also be a pointer to the struct (*NumberFormat)
}

// 通用的map
var docMap map[string]interface{}

func MapStructure() {
	// 将json字符解密成 通用的map[string]interface{}
	docScript := []byte(document)
	json.Unmarshal(docScript, &docMap)
	fmt.Println(docMap)

	// 将map类型转换为 结构体（任意类型）
	var user User
	mapstructure.DecodePath(docMap, &user)
	fmt.Println(user)
}
var document string = "{\n\t\"userContext\": {\n\t\t\"conversationCredentials\": {\n\t            \"sessionToken\": \"06142010_1:75bf6a413327dd71ebe8f3f30c5a4210a9b11e93c028d6e11abfca7ff\"\n\t    },\n\t    \"valid\": true,\n\t    \"isPasswordExpired\": false,\n\t    \"cobrandId\": 10000004,\n\t    \"channelId\": -1,\n\t    \"locale\": \"en_US\",\n\t    \"tncVersion\": 2,\n\t    \"applicationId\": \"17CBE222A42161A3FF450E47CF4C1A00\",\n\t    \"cobrandConversationCredentials\": {\n\t        \"sessionToken\": \"06142010_1:b8d011fefbab8bf1753391b074ffedf9578612d676ed2b7f073b5785b\"\n\t    },\n\t     \"preferenceInfo\": {\n\t         \"currencyCode\": \"USD\",\n\t         \"timeZone\": \"PST\",\n\t         \"dateFormat\": \"MM/dd/yyyy\",\n\t         \"currencyNotationType\": {\n\t             \"currencyNotationType\": \"SYMBOL\"\n\t         },\n\t         \"numberFormat\": {\n\t             \"decimalSeparator\": \".\",\n\t             \"groupingSeparator\": \",\",\n\t             \"groupPattern\": \"###,##0.##\"\n\t         }\n\t     }\n\t },\n\t \"lastLoginTime\": 1375686841,\n\t \"loginCount\": 299,\n\t \"passwordRecovered\": false,\n\t \"emailAddress\": \"johndoe@email.com\",\n\t \"loginName\": \"sptest1\",\n\t \"userId\": 10483860,\n\t \"userType\":\n\t     {\n\t     \"userTypeId\": 1,\n\t     \"userTypeName\": \"normal_user\"\n\t     }\n}"