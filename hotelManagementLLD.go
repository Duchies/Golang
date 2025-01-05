type Hotel struct{
	  Name          string
      ID            int
      Location      Location
}

 const RoomList  []Room

type Location  struct {
	Pin int16
	Street, Area, City, Country string
}

type  Room  struct{
	RoomNumber  string
	RoomStyle  RoomStyle 
	RoomStatus RoomStatus 
	BookingPrice    float32
    RoomKeys        []RoomKey
    HouseKeepingLogs []HouseKeepingLog
}


const (
	StandardRoom RoomStyle = iota 
	DeluxRoom 
	FamilySuite 
)


type RoomStyle int 
type RoomStaus string

const (
	 	AVAILABLE RoomStatus  = "AVAILABLE"
        RESERVED  RoomStatus  = "RESERVED"
        NOT_AVAILABLE RoomStatus  = "NOT_AVAILABLE"
        OCCUPIED RoomStatus  = "OCCUPIED"
        SERVICE_IN_PROGRESS RoomStatus  = "SERVICE_IN_PROGRESS"
)

type RoomKey struct{
	 	KeyId     string
        BarCode   string
        IssuedAt  time.Time
        IsActive  bool
        IsMaster bool
}

type HouseKeepingLog struct {
        Description string
        StartDate   time.Time
        Duration    int
        HouseKeeper *HouseKeeper
}

type HouseKeeper struct {
      Person
}

type Person struct {
       Name       string
       Phone      string
       Account    *Account
}


// Guest represents a hotel guest.
type Guest struct {
        Person
}

type Admin struct{
	Person
}

func (a *Admin) addRoom(roomList []Room){

}
func (a *Admin) deleteRoom(roomNumber string){

}
func (a *Admin) editRoom(room Room){

}

// Receptionist represents a hotel receptionist.
type Receptionist struct {
        Person
        Guest
}

func (r *Receptionist) checkIn(guest Guest,roomDetails BookingDetails){

}

func (r *Receptionist) checkOut(guest Guest){

}


// RoomBooking represents a booking for a room.
type RoomBooking struct {
        BookingID      string
        StartDate     time.Time
        DurationInDays int
        BookingStatus  BookingStatus
        GuestList      []Guest
        RoomInfo       []Room
        TotalRoomCharges BaseRoomCharge
}

// BookingStatus defines the different statuses of a booking.
type BookingStatus int

// BaseRoomCharge is an interface for calculating room charges.
type BaseRoomCharge interface {
        GetCost() float64
}

// RoomCharge represents the base cost of a room.
type RoomCharge struct {
        cost float64
}

// RoomServiceCharge represents additional charges for room service.
type RoomServiceCharge struct {
        cost            float64
        BaseRoomCharge BaseRoomCharge
}

// InRoomPurchaseCharges represents charges for in-room purchases.
type InRoomPurchaseCharges struct {
        cost            float64
        BaseRoomCharge BaseRoomCharge
}


type SearchAndBooking interface {
	SearchRoom() room
	BookingRoom() (bool, string)
}


func (g *Guest)SearchRoom(){

}

func (g *Guest) BookingRoom(){

}


func (g *Receptionist)SearchRoom(){

}

func (g *Receptionist) BookingRoom(){

}




// Account represents an account associated with a person.
type Account struct {
        Username    string
        Password    string
        AccountStatus AccountStatus
}


// AccountStatus defines the different statuses of an account.
type AccountStatus int

const (
        ACTIVE AccountStatus = iota
        CLOSED
        BLOCKED
)


func (room *Room) addRoom() {

} 


func (hk *HouseKeeper) getRoomService(date time.time){

}


