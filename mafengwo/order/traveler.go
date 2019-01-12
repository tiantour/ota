package order

import (
	"encoding/json"

	"github.com/tiantour/ota/mafengwo"
)

type (

	// MFWTravelerItem traveler data
	MFWTravelerItem struct {
		OrderID      string          `json:"order_id"`      // 订单号
		TravelPeople MFWTravelPeople `json:"travel_people"` // 出行人信息
	}
	// MFWTravelPeople traveler people
	MFWTravelPeople struct {
		Traveler  []*MFWTraveler `json:"traveler"`   // 出行人信息
		Trip      MFWTrip        `json:"trip"`       // 出行信息
		TsAddress MFWTsAddress   `json:"ts_address"` // 取还地址
		Address   MFWAddress     `json:"address"`    // 地址信息
	}
	// MFWTraveler traveler
	MFWTraveler struct {
		Name            string `json:"name"`              // 出行人姓名
		IDType          string `json:"id_type"`           // 证件类型 如：当 id_type = 身份证；会返回id_card 目前会返回的类型：身份证/护照/港澳通行证/台湾通行证
		Birthday        string `json:"birthday"`          // 出生年月日
		Gender          string `json:"gender"`            // 性别
		Nationality     string `json:"nationality"`       // 国籍
		Height          int32  `json:"height"`            // 身高(cm)
		Weight          int32  `json:"weight"`            // 体重(kg)
		ShoeSize        int32  `json:"shoe_size"`         // 鞋码(欧码)
		LeftEyeSight    int32  `json:"left_eye_sight"`    // 左眼视力
		RightEyeSight   int32  `json:"right_eye_sight"`   // 右眼视力
		DateOfEexpiry   string `json:"date_of_expiry"`    // 有效期
		Cellphone       string `json:"cellphone"`         // 联系电话
		FamilyName      string `json:"family_name"`       // 姓（拼音）
		MainlandPhone   string `json:"mainland_phone"`    // 境内手机号
		TravelerID      int32  `json:"traveler_id"`       // 出行人ID
		LaissezPasserTW string `json:"laissez_passer_tw"` // 台湾通行证
		LaissezPasser   string `json:"laissez_passer"`    // 港澳通行证
		Passport        string `json:"passport"`          // 护照证件号
		IDCard          string `json:"id_card"`           // 身份证号码
		FirstName       string `json:"first_name"`        // 名（拼音）
	}
	// MFWTrip trip
	MFWTrip struct {
		PickUpTime            string `json:"pick_up_time"`              // 接人时间
		PickUpPlace           string `json:"pick_up_place"`             // 接人地点
		SendTo                string `json:"send_to"`                   // 送达地点
		PickUpPlaceEN         string `json:"pick_up_place_en"`          // 接人地点(英文)
		SendToEN              string `json:"send_to_en"`                // 送达地点(英文)
		HotelNamePickUp       string `json:"hotel_name_pick_up"`        // 接人酒店名称
		HotelAddressPickUp    string `json:"hotel_address_pick_up"`     // 接人酒店地址
		HotelNameENPickUp     string `json:"hotel_name_en_pick_up"`     // 接人酒店名称(英文)
		HotelAddressENPickUp  string `json:"hotel_address_en_pick_up"`  // 接人酒店地址(英文)
		HotelNameSendTo       string `json:"hotel_name_send_to"`        // 送达酒店名称
		HotelAddressSendTo    string `json:"hotel_address_send_to"`     // 送达酒店地址
		HotelNameENSendTo     string `json:"hotel_name_en_send_to"`     // 送达酒店名称(英文)
		HotelAddressENSendTo  string `json:"hotel_address_en_send_to"`  // 送达酒店地址（英文）
		HotelNameOverNight    string `json:"hotel_name_over_night"`     // 过夜酒店名称(英文)
		FlightNoArrival       string `json:"flight_no_arrival"`         // 接机航班号
		FlightTimeArrival     string `json:"flight_time_arrival"`       // 航班抵达时间
		FlightNoDeparture     string `json:"flight_no_departure"`       // 送机航班号
		FlightTimeDeparture   string `json:"flight_time_departure"`     // 航班起飞时间
		Luggage               string `json:"luggage"`                   // 行李数
		HotelNameEN           string `json:"hotel_name_en"`             // 接送酒店名称(英文)
		HotelDddressEN        string `json:"hotel_address_en"`          // 接送酒店地址(英文)
		HotelPhone            string `json:"hotel_phone"`               // 接人酒店电话
		HotelTelephone        string `json:"hotel_telephone"`           // 送达酒店电话
		UsingTime             string `json:"using_time"`                // 用车时间
		UsingPlace            string `json:"using_place"`               // 用车地点
		FlightNumber          string `json:"flight_number"`             // 接机/送机航班号
		FlightTime            string `json:"flight_time"`               // 降落/起飞时间
		Place                 string `json:"place"`                     // 接人/送达地点
		ReturnHotel           string `json:"return_hotel"`              // 返回酒店名称(英文)
		HotelAdress           string `json:"hotel_adress"`              // 返回酒店地址(英文)
		ReturnHotelPhone      string `json:"return_hotel_phone"`        // 返回酒店电话
		PickAndSendHotelPhone string `json:"pick_and_send_hotel_phone"` // 接送酒店电话
		Schedule              string `json:"schedule"`                  // 行程计划
		BackHotel             string `json:"back_hotel"`                // 返程送回酒店名称（英文）
		BackAdress            string `json:"back_address"`              // 返程送回酒店地址（英文）
		BackPhone             string `json:"back_phone"`                // 返程送回酒店电话
		HotelName             string `json:"hotel_name"`                // 酒店名称（英文）
		HotelAddress          string `json:"hotel_address"`             // 酒店地址（英文）
		HotelPhoneNumber      string `json:"hotel_phone_number"`        // 酒店联系电话
		CheckInDate           string `json:"check_in_date"`             // 入住酒店日期
		CheckOutDate          string `json:"check_out_date"`            // 离开酒店日期
		ReturnFlightNumber    string `json:"return_flight_number"`      // 返程航班号/火车列次
		ReturnFlightTime      string `json:"return_flight_time"`        // 返程航班/列次时间
		ArrivalDate           string `json:"arrival_date"`              // 航班抵达日期
		DepartureDate         string `json:"departure_date"`            // 航班起飞日期
		DepartureHotelName    string `json:"departure_hotel_name"`      // 出发酒店名称(英文)
		DepartureHotelAdress  string `json:"departure_hotel_adress"`    // 出发酒店地址(英文)
		DepartureHotelNumber  string `json:"departure_hotel_number"`    // 出发酒店电话
		BackDate              string `json:"back_date"`                 // 返程日期
		OverNightHotelAddress string `json:"over_night_hotel_address"`  // 过夜酒店地址(英文)
		GetDeviceAdress       string `json:"get_device_adress"`         // 取还设备地址
		DepartureHotelNameCN  string `json:"departure_hotel_name_cn"`   // 出发酒店名称（中文）
		Time                  string `json:"time"`                      // 预约时间
		Number                string `json:"number"`                    // 用车人数
		Phone                 string `json:"phone"`                     // 境外联系电话
		Wechat                string `json:"wechat"`                    // 微信号
		EstimatedTravelDate   string `json:"estimated_travel_date"`     // 预计出行日期
		TrainNumber           string `json:"train_number"`              // 到达火车车次
		TrainStation          string `json:"train_station"`             // 到达车站
		DepartureTime         string `json:"departure_time"`            // 出发时间(单选)
		DepartureFrequency    string `json:"departure_frequency"`       // 出发班次
		DepartureHotelArea    string `json:"departure_hotel_area"`      // 出发酒店所在区域
		MealTime              string `json:"meal_time"`                 // 用餐时间
		TouristsNumber        string `json:"tourists_number"`           // 出行人个数
	}
	// MFWTsAddress ts_address
	MFWTsAddress struct {
		PickUpAddress string `json:"pick_up_address"` // 取件地址
		ReturnAddress string `json:"return_address"`  // 还件地址
	}
	// MFWAddress address
	MFWAddress struct {
		Adress        string `json:"adress"`         // 收货地址
		ReceiverName  string `json:"receiver_name"`  // 收件人姓名
		ReceiverPhone string `json:"receiver_phone"` // 收货人电话
	}
	// MFWTravelerParam traveler param
	MFWTravelerParam struct {
		OrderID string          `json:"order_id"` // 订单号
		Params  MFWTravelPeople `json:"params"`   // 出行人信息
	}
)

// Traveler traveler
type Traveler struct {
	OrderID    string           `json:"order_id"`    // 订单号
	TravelerID int32            `json:"traveler_id"` // 出行人ID
	Params     MFWTravelerParam `json:"params"`      // 出行人信息
}

// NewTraveler new traveler
func NewTraveler() *Traveler {
	return &Traveler{}
}

// List get traveler list
func (t *Traveler) List(orderID string) (*MFWTravelerItem, error) {
	action := "sales.order.traveler.get"
	data, err := json.Marshal(&Traveler{
		OrderID: orderID,
	})
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := MFWTravelerItem{}
	err = json.Unmarshal(body, &result)
	return &result, err
}

// Update update traveler
func (t *Traveler) Update(args *Traveler) (*Traveler, error) {
	action := "sales.order.traveler.update"
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := mafengwo.NewDeals().Fetch(action, data)
	if err != nil {
		return nil, err
	}
	result := Traveler{}
	err = json.Unmarshal(body, &result)
	return &result, err
}
