package powerxtypes

type GetHomeReply struct {
	Greet       string `json:"greet,optional"`
	Description string `json:"description,optional"`
	Version     string `json:"version,optional"`
}

type GetEmployeeOptionsRequest struct {
	LikeName        string `form:"likeName,optional"`
	LikeEmail       string `form:"likeEmail,optional"`
	LikePhoneNumber string `form:"likePhoneNumber,optional"`
	PageIndex       int    `form:"pageIndex,optional"`
	PageSize        int    `form:"pageSize,optional"`
}

type EmployeeOption struct {
	Id          int64  `json:"id"`
	Avatar      string `json:"avatar"`
	Account     string `json:"account"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type GetEmployeeOptionsReply struct {
	List      []EmployeeOption `json:"list"`
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Total     int64            `json:"total"`
}

type EmployeeQueryRoleOption struct {
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}

type EmployeeQueryDepartmentOption struct {
	DepartmentId   int64  `json:"departmentId"`
	DepartmentName string `json:"departmentName"`
}

type GetEmployeeQueryOptionsReply struct {
	Positions   []string                        `json:"positions"`
	Roles       []EmployeeQueryRoleOption       `json:"roles"`
	Departments []EmployeeQueryDepartmentOption `json:"departments"`
}

type GetDepartmentOptionsRequest struct {
	Ids       []int64 `form:"ids,optional"`
	LikeName  string  `form:"likeName,optional"`
	PageIndex int     `form:"pageIndex,optional"`
	PageSize  int     `form:"pageSize,optional"`
}

type DepartmentOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GetDepartmentOptionsReply struct {
	List      []DepartmentOption `json:"list"`
	PageIndex int                `json:"pageIndex"`
	PageSize  int                `json:"pageSize"`
	Total     int64              `json:"total"`
}

type DepartmentLeader struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

type DepartmentNode struct {
	Id          int64            `json:"id"`
	DepName     string           `json:"depName"`
	Leader      DepartmentLeader `json:"leader"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Remark      string           `json:"remark"`
	Children    []DepartmentNode `json:"children"`
}

type GetDepartmentTreeRequest struct {
	DepId int64 `path:"depId"`
}

type GetDepartmentTreeReply struct {
	DepTree DepartmentNode `json:"depTree"`
}

type CreateDepartmentRequest struct {
	DepName     string `json:"depName"`
	LeaderId    int64  `json:"leaderId"`
	PId         int64  `json:"pId"`
	Desc        string `json:"desc,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Remark      string `json:"remark,optional"`
}

type CreateDepartmentReply struct {
	Id int64 `json:"id"`
}

type DeleteDepartmentRequest struct {
	Id int64 `path:"id"`
}

type DeleteDepartmentReply struct {
	Id int64 `json:"id"`
}

type PatchDepartmentRequest struct {
	Id          int64  `path:"id"`
	DepName     string `json:"depName,optional"`
	LeaderId    int64  `json:"leaderId,optional"`
	PId         int64  `json:"pId,optional"`
	Desc        string `json:"desc,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Remark      string `json:"remark,optional"`
}

type PatchDepartmentReply struct {
	*Department
}

type Department struct {
	Id          int64            `json:"id"`
	DepName     string           `json:"depName"`
	Leader      DepartmentLeader `json:"leader"`
	PhoneNumber string           `json:"phoneNumber"`
	Email       string           `json:"email"`
	Remark      string           `json:"remark"`
}

type GetDepartmentRequest struct {
	Id int64 `path:"id"`
}

type GetDepartmentReply struct {
	*Department
}

type GetEmployeeRequest struct {
	Id int64 `path:"id"`
}

type GetEmployeeReply struct {
	*Employee
}

type ListEmployeesRequest struct {
	Ids             []int64  `form:"ids,optional"`
	LikeName        string   `form:"likeName,optional"`
	LikeEmail       string   `form:"likeEmail,optional"`
	DepIds          []int64  `form:"depIds,optional"`
	Positions       []string `form:"positions,optional"`
	LikePhoneNumber string   `form:"likePhoneNumber,optional"`
	RoleCodes       []string `form:"roleCodes,optional"`
	IsEnabled       *bool    `form:"isEnable,optional"`
	PageIndex       int      `form:"pageIndex,optional"`
	PageSize        int      `form:"pageSize,optional"`
}

type EmployeeDepartment struct {
	DepId   int64  `json:"depId"`
	DepName string `json:"depName"`
}

type Employee struct {
	Id            int64               `json:"id"`
	Account       string              `json:"account"`
	Name          string              `json:"name"`
	Email         string              `json:"email"`
	MobilePhone   string              `json:"mobilePhone"`
	Gender        string              `json:"gender"`
	NickName      string              `json:"nickName,optional"`
	Desc          string              `json:"desc,optional"`
	Avatar        string              `json:"avatar,optional"`
	ExternalEmail string              `json:"externalEmail,optional"`
	Roles         []string            `json:"roles"`
	Department    *EmployeeDepartment `json:"department"`
	Position      string              `json:"position"`
	JobTitle      string              `json:"jobTitle"`
	IsEnabled     bool                `json:"isEnabled"`
	CreatedAt     string              `json:"createdAt"`
}

type ListEmployeesReply struct {
	List      []Employee `json:"list"`
	PageIndex int        `json:"pageIndex"`
	PageSize  int        `json:"pageSize"`
	Total     int64      `json:"total"`
}

type SyncEmployeesRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type SyncEmployeesReply struct {
	Status bool `json:"status"`
}

type CreateEmployeeRequest struct {
	Account       string `json:"account"`
	Name          string `json:"name"`
	NickName      string `json:"nickName,optional"`
	Desc          string `json:"desc,optional"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar,optional"`
	ExternalEmail string `json:"externalEmail,optional"`
	MobilePhone   string `json:"mobilePhone,optional"`
	Gender        string `json:"gender,options=male|female|un_know"`
	DepId         int64  `json:"depId"`
	Position      string `json:"position,optional"`
	JobTitle      string `json:"jobTitle,optional"`
	Password      string `json:"password,optional"`
}

type CreateEmployeeReply struct {
	Id int64 `json:"id"`
}

type UpdateEmployeeRequest struct {
	Id            int64  `path:"id"`
	Name          string `json:"name,optional"`
	NickName      string `json:"nickName,optional"`
	Desc          string `json:"desc,optional"`
	Email         string `json:"email,optional"`
	Avatar        string `json:"avatar,optional"`
	ExternalEmail string `json:"externalEmail,optional"`
	MobilePhone   string `json:"mobilePhone,optional"`
	Gender        string `json:"gender,optional,options=male|female|un_know"`
	DepId         int64  `json:"depId,optional"`
	Position      string `json:"position,optional"`
	JobTitle      string `json:"jobTitle,optional"`
	Password      string `json:"password,optional"`
	Status        string `json:"status,optional,options=enabled|disabled"`
}

type UpdateEmployeeReply struct {
	*Employee
}

type DeleteEmployeeRequest struct {
	Id int64 `path:"id"`
}

type DeleteEmployeeReply struct {
	Id int64 `json:"id"`
}

type ResetPasswordRequest struct {
	UserId int64 `json:"userId"`
}

type ResetPasswordReply struct {
	Status string `json:"status"`
}

type AdminAPI struct {
	Id        int64  `json:"id"`
	API       string `json:"api"`
	Method    string `json:"method"`
	Name      string `json:"name"`
	GroupId   int64  `json:"groupId"`
	GroupName string `json:"groupName"`
	Desc      string `json:"desc"`
}

type AdminRole struct {
	RoleCode   string     `json:"roleCode"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc"`
	IsReserved bool       `json:"isReserved"`
	APIList    []AdminAPI `json:"apiList"`
	MenuNames  []string   `json:"menuNames"`
}

type ListRolesReply struct {
	List []AdminRole `json:"list"`
}

type CreateRoleRequest struct {
	RoleCode  string   `json:"roleCode"`
	Name      string   `json:"name"`
	Desc      string   `json:"desc"`
	APIIds    []int64  `json:"apiIds"`
	MenuNames []string `json:"menuNames"`
}

type CreateRoleReply struct {
	RoleCode string `json:"roleCode"`
}

type GetRoleRequest struct {
	RoleCode string `path:"roleCode"`
}

type GetRoleReply struct {
	*AdminRole
}

type PatchRoleReqeust struct {
	RoleCode  string   `path:"roleCode"`
	Name      string   `json:"name"`
	Desc      string   `json:"desc,optional"`
	APIIds    []int64  `json:"apiIds,optional"`
	MenuNames []string `json:"menuNames,optional"`
}

type PatchRoleReply struct {
	*AdminRole
}

type SetRolePermissionsRequest struct {
	RoleCode string  `path:"roleCode"`
	APIIds   []int64 `json:"apiIds"`
}

type SetRolePermissionsReply struct {
	Status string `json:"status"`
}

type SetRoleEmployeesRequest struct {
	RoleCode    string  `path:"roleCode"`
	EmployeeIds []int64 `json:"employeeIds"`
}

type SetRoleEmployeesReply struct {
	Status string `json:"status"`
}

type ListAPIRequest struct {
	GroupId int64 `form:"groupId,optional"`
}

type ListAPIReply struct {
	List []AdminAPI `json:"list"`
}

type GetRoleEmployeesReqeust struct {
	RoleCode  string `path:"roleCode"`
	PageIndex int    `form:"pageIndex"`
	PageSize  int    `form:"pageSize"`
}

type RoleEmployeeDepartment struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type RoleEmployee struct {
	Id          int64                   `json:"id"`
	Name        string                  `json:"name"`
	Nickname    string                  `json:"nickname"`
	Account     string                  `json:"account"`
	PhoneNumber string                  `json:"phoneNumber"`
	Department  *RoleEmployeeDepartment `json:"department"`
	Email       string                  `json:"email"`
}

type GetRoleEmployeesReply struct {
	List      []RoleEmployee `json:"list"`
	PageIndex int            `json:"pageIndex"`
	PageSize  int            `json:"pageSize"`
	Total     int64          `json:"total"`
}

type SetUserRolesRequest struct {
	UserId    int64    `path:"userId"`
	RoleCodes []string `json:"roleCodes"`
}

type SetUserRolesReply struct {
	Status string `json:"status"`
}

type LoginRequest struct {
	UserName    string `json:"userName,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Email       string `json:"email,optional"`
	Password    string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ExchangeRequest struct {
	Type string `path:"type,optional=wechat"`
	Code string `json:"code"`
}

type ExchangeReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ListDictionaryTypesPageRequest struct {
	PageIndex int `form:"pageIndex,optional"`
	PageSize  int `form:"pageSize,optional"`
}

type DictionaryType struct {
	Id          int64             `json:"id,optional"`
	Type        string            `json:"type"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Items       []*DictionaryItem `json:"items,optional"`
}

type ListDictionaryTypesPageReply struct {
	List      []DictionaryType `json:"list"`
	PageIndex int              `json:"pageIndex"`
	PageSize  int              `json:"pageSize"`
	Total     int64            `json:"total"`
}

type GetDictionaryTypeRequest struct {
	DictionaryType string `path:"type"`
}

type GetDictionaryTypeReply struct {
	*DictionaryType
}

type CreateDictionaryTypeRequest struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,optional"`
}

type CreateDictionaryTypeReply struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description,optional"`
}

type UpdateDictionaryTypeRequest struct {
	Type        string `path:"type"`
	Name        string `json:"name,optional"`
	Description string `json:"description,optional"`
}

type UpdateDictionaryTypeReply struct {
	*DictionaryType
}

type DeleteDictionaryTypeRequest struct {
	Type string `path:"type"`
}

type DeleteDictionaryTypeReply struct {
	Type string `json:"type"`
}

type ListDictionaryItemsRequest struct {
	Type      string `form:"type"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type DictionaryItem struct {
	Id          int64  `json:"id,optional"`
	Key         string `json:"key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Sort        int    `json:"sort"`
	Description string `json:"description"`
}

type ListDictionaryItemsReply struct {
	List []DictionaryItem `json:"list"`
}

type GetDictionaryItemRequest struct {
	DictionaryType string `path:"type"`
	DictionaryItem string `path:"key"`
}

type GetDictionaryItemReply struct {
	*DictionaryItem
}

type CreateDictionaryItemRequest struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Sort        int    `json:"sort,optional"`
	Description string `json:"description,optional"`
}

type CreateDictionaryItemReply struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Name        string `json:"name,optional"`
	Value       string `json:"value,optional"`
	Sort        int    `json:"sort,optional"`
	Description string `json:"description,optional"`
}

type UpdateDictionaryItemRequest struct {
	Key         string `path:"key"`
	Type        string `path:"type"`
	Name        string `json:"name,optional"`
	Value       string `json:"value,optional"`
	Sort        int    `json:"sort,optional"`
	Description string `json:"description,optional"`
}

type UpdateDictionaryItemReply struct {
	*DictionaryItem
}

type DeleteDictionaryItemRequest struct {
	Key  string `path:"key"`
	Type string `path:"type"`
}

type DeleteDictionaryItemReply struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

type GetUserInfoReply struct {
	Id            int64    `json:"id"`
	Account       string   `json:"account"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	MobilePhone   string   `json:"mobilePhone"`
	Gender        string   `json:"gender"`
	NickName      string   `json:"nickName"`
	Desc          string   `json:"desc"`
	Avatar        string   `json:"avatar"`
	ExternalEmail string   `json:"externalEmail"`
	DepName       string   `json:"depName"`
	Position      string   `json:"position"`
	JobTitle      string   `json:"jobTitle"`
	IsEnabled     bool     `json:"isEnabled"`
	CreatedAt     string   `json:"createdAt"`
	Roles         []string `json:"roles"`
}

type MenuRoles struct {
	MenuName       string   `json:"menuName"`
	AllowRoleCodes []string `json:"allowRoleCodes"`
}

type GetMenuRolesReply struct {
	MenuRoles []MenuRoles `json:"menuRoles"`
}

type ModifyPasswordReqeust struct {
	Password string `json:"password"`
}

type MediaResource struct {
	Id            int64  `json:"id,optional"`
	Filename      string `json:"filename,optional"`
	Size          int64  `json:"size,optional"`
	Url           string `json:"url,optional"`
	BucketName    string `json:"bucketName,optional"`
	IsLocalStored bool   `json:"isLocalStored,optional"`
	ContentType   string `json:"contentType,optional"`
	ResourceType  string `json:"resourceType,optional"`
}

type CreateMediaResourceReply struct {
	*MediaResource
	IsOSS bool `json:"isOSS"`
}

type GetMediaResourceRequest struct {
	Id int64 `path:"id"`
}

type GetMediaResourceReply struct {
	MediaResource *MediaResource `json:"mediaResource"`
}

type ListMediaResourcesPageRequest struct {
	LikeName  string `form:"likeName,optional"`
	OrderBy   string `form:"orderBy,optional"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type ListMediaResourcesPageReply struct {
	List      []MediaResource `json:"list,optional"`
	PageIndex int             `json:"pageIndex,optional"`
	PageSize  int             `json:"pageSize,optional"`
	Total     int64           `json:"total,optional"`
}

type DeleteMediaResourceRequest struct {
	Id int64 `path:"id"`
}

type DeleteMediaResourceReply struct {
	MediaResourceId int64 `json:"id"`
}

type LeadExternalId struct {
	OpenIdInMiniProgram           string `json:"openIdInMiniProgram,optional"`
	OpenIdInWeChatOfficialAccount string `json:"openIdInWeChatOfficialAccount,optional"`
	OpenIdInWeCom                 string `json:"openIdInWeCom,optional"`
}

type LeadInviter struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Mobile string `json:"mobile,optional"`
	Email  string `json:"email,optional"`
}

type Lead struct {
	Id          int64        `json:"id,optional"`
	Name        string       `json:"name"`
	Mobile      string       `json:"mobile"`
	Email       string       `json:"email,optional"`
	Inviter     *LeadInviter `json:"inviter,optional"`
	InviterId   int64        `json:"inviter,optional"`
	Source      int          `json:"source,optional"`
	Type        int          `json:"type,optional"`
	IsActivated bool         `json:"isActivated,optional,omitempty"`
	CreatedAt   string       `json:"createdAt,optional"`
	*LeadExternalId
}

type GetLeadReqeuest struct {
	Id int64 `path:"id"`
}

type GetLeadReply struct {
	Lead *Lead `json:"lead"`
}

type ListLeadsPageRequest struct {
	LikeName   string `form:"likeName,optional"`
	LikeMobile string `form:"likeMobile,optional"`
	Sources    []int  `form:"sources,optional"`
	Statuses   []int  `form:"statuses,optional"`
	OrderBy    string `form:"orderBy,optional"`
	PageIndex  int    `form:"pageIndex,optional"`
	PageSize   int    `form:"pageSize,optional"`
}

type ListLeadsPageReply struct {
	List      []Lead `json:"list,optional"`
	PageIndex int    `json:"pageIndex,optional"`
	PageSize  int    `json:"pageSize,optional"`
	Total     int64  `json:"total,optional"`
}

type CreateLeadRequest struct {
	Lead
}

type CreateLeadReply struct {
	LeadId int64 `json:"id"`
}

type PutLeadRequest struct {
	LeadId int64 `path:"id"`
	Lead
}

type PutLeadReply struct {
	*Lead
}

type PatchLeadRequest struct {
	LeadId      int64  `path:"id"`
	Name        string `json:"name,optional"`
	Email       string `json:"email,optional"`
	InviterId   int64  `json:"inviterId,optional"`
	Source      int    `json:"source,optional"`
	Type        int    `json:"type,optional"`
	IsActivated bool   `json:"isActivated,optional,omitempty"`
}

type PatchLeadReply struct {
	*Lead
}

type DeleteLeadRequest struct {
	Id int64 `path:"id"`
}

type DeleteLeadReply struct {
	LeadId int64 `json:"id"`
}

type AssignLeadToEmployeeRequest struct {
	Id         string `path:"id"`
	EmployeeId int64  `json:"employeeId"`
}

type AssignLeadToEmployeeReply struct {
	LeadId int64 `json:"leadId"`
}

type CustomerExternalId struct {
	OpenIdInMiniProgram           string `json:"openIdInMiniProgram,optional"`
	OpenIdInWeChatOfficialAccount string `json:"openIdInWeChatOfficialAccount,optional"`
	OpenIdInWeCom                 string `json:"openIdInWeCom,optional"`
}

type CustomerInviter struct {
	Id     int64  `json:"id"`
	Name   string `json:"name,optional"`
	Mobile string `json:"mobile,optional"`
	Email  string `json:"email,optional"`
}

type Customer struct {
	Id          int64            `json:"id,optional"`
	Name        string           `json:"name"`
	Mobile      string           `json:"mobile"`
	Email       string           `json:"email,optional"`
	Inviter     *CustomerInviter `json:"inviter,optional"`
	InviterId   int64            `json:"inviter,optional"`
	Source      int              `json:"source,optional"`
	Type        int              `json:"type,optional"`
	IsActivated bool             `json:"isActivated,optional,omitempty"`
	CreatedAt   string           `json:"createdAt,optional"`
	*CustomerExternalId
}

type GetCustomerReqeuest struct {
	Id int64 `path:"id"`
}

type GetCustomerReply struct {
	Customer *Customer `json:"customerdomain"`
}

type ListCustomersPageRequest struct {
	LikeName   string `form:"likeName,optional"`
	LikeMobile string `form:"likeMobile,optional"`
	Sources    []int  `form:"sources,optional"`
	Statuses   []int  `form:"statuses,optional"`
	OrderBy    string `form:"orderBy,optional"`
	PageIndex  int    `form:"pageIndex,optional"`
	PageSize   int    `form:"pageSize,optional"`
}

type ListCustomersPageReply struct {
	List      []Customer `json:"list,optional"`
	PageIndex int        `json:"pageIndex,optional"`
	PageSize  int        `json:"pageSize,optional"`
	Total     int64      `json:"total,optional"`
}

type CreateCustomerRequest struct {
	Customer
}

type CreateCustomerReply struct {
	CustomerId int64 `json:"id"`
}

type PutCustomerRequest struct {
	CustomerId int64 `path:"id"`
	Customer
}

type PutCustomerReply struct {
	*Customer
}

type PatchCustomerRequest struct {
	CustomerId  int64  `path:"id"`
	Name        string `json:"name,optional"`
	Email       string `json:"email,optional"`
	InviterId   int64  `json:"inviterId,optional"`
	Source      int    `json:"source,optional"`
	Type        int    `json:"type,optional"`
	IsActivated bool   `json:"isActivated,optional,omitempty"`
}

type PatchCustomerReply struct {
	*Customer
}

type DeleteCustomerRequest struct {
	Id int64 `path:"id"`
}

type DeleteCustomerReply struct {
	CustomerId int64 `json:"id"`
}

type AssignCustomerToEmployeeRequest struct {
	Id         string `path:"id"`
	EmployeeId int64  `json:"employeeId"`
}

type AssignCustomerToEmployeeReply struct {
	CustomerId int64 `json:"customerId"`
}

type ListMediasPageRequest struct {
	MediaTypes []int8   `form:"mediaTypes,optional"`
	Keys       []string `form:"keys,optional"`
	OrderBy    string   `form:"orderBy,optional"`
	PageIndex  int      `form:"pageIndex,optional"`
	PageSize   int      `form:"pageSize,optional"`
}

type MediaImage struct {
	Id            int64  `json:"id,optional"`
	Filename      string `json:"filename,optional"`
	Size          int64  `json:"size,optional"`
	Url           string `json:"url,optional"`
	BucketName    string `json:"bucketName,optional"`
	IsLocalStored bool   `json:"isLocalStored,optional"`
	ContentType   string `json:"contentType,optional"`
	ResourceType  string `json:"resourceType,optional"`
}

type Media struct {
	Id             int64         `json:"id,optional"`
	Title          string        `json:"title,optional"`
	SubTitle       string        `json:"subTitle,optional"`
	CoverImageId   int64         `json:"coverImageId,optional"`
	ResourceUrl    string        `json:"resourceUrl,optional"`
	Description    string        `json:"description,optional"`
	MediaType      int           `json:"mediaType,optional"`
	ViewedCount    int           `json:"viewedCount,optional"`
	CoverImage     *MediaImage   `json:"coverImage,optional"`
	DetailImageIds []int64       `json:"detailImageIds,optional"`
	DetailImages   []*MediaImage `json:"detailImages,optional"`
}

type ListMediasPageReply struct {
	List      []*Media `json:"list"`
	PageIndex int      `json:"pageIndex"`
	PageSize  int      `json:"pageSize"`
	Total     int64    `json:"total"`
}

type CreateMediaRequest struct {
	Media
}

type CreateMediaReply struct {
	MediaId int64 `json:"id"`
}

type UpdateMediaRequest struct {
	MediaId int64 `path:"id"`
	Media
}

type UpdateMediaReply struct {
	MediaId int64 `json:"id"`
}

type GetMediaRequest struct {
	MediaId int64 `path:"id"`
}

type GetMediaReply struct {
	*Media
}

type DeleteMediaRequest struct {
	MediaId int64 `path:"id"`
}

type DeleteMediaReply struct {
	MediaId int64 `json:"id"`
}

type StoreArtisanSpecific struct {
	ArtisanId int64 `json:"artisanId,optional"`
}

type StoreImage struct {
	Id            int64  `json:"id,optional"`
	Filename      string `json:"filename,optional"`
	Size          int64  `json:"size,optional"`
	Url           string `json:"url,optional"`
	BucketName    string `json:"bucketName,optional"`
	IsLocalStored bool   `json:"isLocalStored,optional"`
	ContentType   string `json:"contentType,optional"`
	ResourceType  string `json:"resourceType,optional"`
}

type StoreArtisan struct {
	EmployeeId      int64                `json:"employeeId,optional"`
	Name            string               `json:"name,optional"`
	Level           int8                 `json:"level,optional"`
	Gender          bool                 `json:"gender,optional"`
	Birthday        string               `json:"birthday,optional"`
	PhoneNumber     string               `json:"phoneNumber,optional"`
	CoverURL        string               `json:"coverURL,optional"`
	WorkNo          string               `json:"workNo,optional"`
	Email           string               `json:"email,optional"`
	Experience      string               `json:"experience,optional"`
	Specialty       string               `json:"specialty,optional"`
	Certificate     string               `json:"certificate,optional"`
	Address         string               `json:"address,optional"`
	ArtisanSpecific StoreArtisanSpecific `json:"artisanSpecific,optional"`
}

type Store struct {
	Id              int64           `json:"id,optional"`
	Name            string          `json:"name"`
	StoreEmployeeId int64           `json:"storeEmployeeId,optional"`
	ContactNumber   string          `json:"contactNumber"`
	Email           string          `json:"email,optional"`
	Address         string          `json:"address"`
	Description     string          `json:"description,optional"`
	Longitude       float32         `json:"longitude,optional"`
	Latitude        float32         `json:"latitude,optional"`
	StartWork       string          `json:"startWork,optional"`
	EndWork         string          `json:"endWork,optional"`
	Artisans        []*StoreArtisan `json:"artisans,optional"`
	CreatedAt       string          `json:"createdAt,optional"`
	CoverImageId    int64           `json:"coverImageId,optional"`
	CoverImage      *StoreImage     `json:"coverImage,optional"`
	DetailImageIds  []int64         `json:"detailImageIds,optional"`
	DetailImages    []*StoreImage   `json:"detailImages,optional"`
}

type ListStoresPageRequest struct {
	Ids       []int64 `form:"ids,optional"`
	LikeName  string  `form:"likeName,optional"`
	OrderBy   string  `form:"orderBy,optional"`
	PageIndex int     `form:"pageIndex,optional"`
	PageSize  int     `form:"pageSize,optional"`
}

type ListStoresPageReply struct {
	List      []*Store `json:"list"`
	PageIndex int      `json:"pageIndex"`
	PageSize  int      `json:"pageSize"`
	Total     int64    `json:"total"`
}

type CreateStoreRequest struct {
	Store
}

type CreateStoreReply struct {
	StoreId int64 `json:"id"`
}

type GetStoreRequest struct {
	StoreId int64 `path:"id"`
}

type GetStoreReply struct {
	*Store
}

type PutStoreRequest struct {
	StoreId int64 `path:"id"`
	Store
}

type PutStoreReply struct {
	*Store
}

type DeleteStoreRequest struct {
	StoreId int64 `path:"id"`
}

type DeleteStoreReply struct {
	StoreId int64 `json:"id"`
}

type AssignStoreManagerRequest struct {
	Id         int64 `path:"id"`
	EmployeeId int64 `json:"employeeId"`
}

type AssignStoreManagerReply struct {
	Store
}

type GetOpportunityListRequest struct {
	Name      string `form:"name,optional"`
	Source    string `form:"source,optional"`
	Type      string `form:"type,optional"`
	Stage     string `form:"stage,optional"`
	OrderBy   string `form:"orderBy,optional"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type Opportunity struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Requirement string  `json:"requirement"`
	CustomerId  int64   `json:"customerId"`
	Probability float32 `json:"probability"`
	Source      string  `json:"source"`
	Type        string  `json:"type"`
	EmployeeId  int64   `json:"employeeId"`
	Stage       string  `json:"stage"`
	ClosedDate  string  `json:"closedDate"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

type GetOpportunityListReply struct {
	List      []Opportunity `json:"list"`
	PageIndex int           `json:"pageIndex"`
	PageSize  int           `json:"pageSize"`
	Total     int64         `json:"total"`
}

type CreateOpportunityRequest struct {
	Name        string  `json:"name"`
	Requirement string  `json:"requirement"`
	CustomerId  int64   `json:"customerId"`
	Probability float32 `json:"probability,optional"`
	Source      string  `json:"source,options=new_customer|old_customer_new_purchase|old_customer_repurchase|old_customer_upgrade"`
	Type        string  `json:"type,options=trial_requirement|requirement_match|detailed_requirement_analysis|solution_provided|quotation|negotiation|closed_unsuccessful|closed_successful"`
	EmployeeId  int64   `json:"employeeId"`
	Stage       string  `json:"stage"`
}

type CreateOpportunityReply struct {
	Id int64 `json:"id"`
}

type AssignEmployeeToOpportunityRequest struct {
	Id         int64 `path:"id"`
	EmployeeId int64 `json:"employeeId"`
}

type AssignEmployeeToOpportunityReply struct {
	Id int64 `json:"id"`
}

type UpdateOpportunityRequest struct {
	Id          int64   `path:"id"`
	Name        string  `json:"name,optional"`
	Requirement string  `json:"requirement,optional"`
	CustomerId  int64   `json:"customerId,optional"`
	Probability float32 `json:"probability,optional"`
	Source      string  `json:"source,optional,options=new_customer|old_customer_new_purchase|old_customer_repurchase|old_customer_upgrade"`
	Type        string  `json:"type,optional,options=trial_requirement|requirement_match|detailed_requirement_analysis|solution_provided|quotation|negotiation|closed_unsuccessful|closed_successful"`
	EmployeeId  int64   `json:"employeeId,optional"`
	Stage       string  `json:"stage,optional"`
	ClosedDate  string  `json:"closedDate,optional"`
}

type UpdateOpportunityReply struct {
	*Opportunity
}

type DeleteOpportunityRequest struct {
	Id int64 `path:"id"`
}

type DeleteOpportunityReply struct {
	Id int64 `json:"id"`
}

type PriceBook struct {
	Id          int64  `json:"id,optional"`
	IsStandard  bool   `json:"isStandard,optional"`
	Name        string `json:"name"`
	Description string `json:"description,optional"`
	StoreId     int64  `json:"storeId,optional"`
	CreatedAt   string `json:"createdAt,optional"`
}

type ListPriceBooksPageRequest struct {
	LikeName  string `json:"likeName,optional"`
	StoreId   int64  `json:"storeId,optional"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type ListPriceBooksPageReply struct {
	List      []PriceBook `json:"list"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Total     int64       `json:"total"`
}

type GetPriceBookRequest struct {
	PriceBook int64 `path:"id"`
}

type GetPriceBookReply struct {
	*PriceBook
}

type UpsertPriceBookRequest struct {
	PriceBook
}

type UpsertPriceBookReply struct {
	*PriceBook
}

type DeletePriceBookRequest struct {
	Id int64 `path:"id"`
}

type DeletePriceBookReply struct {
	Id int64 `json:"id"`
}

type ProductImage struct {
	Id            int64  `json:"id,optional"`
	Filename      string `json:"filename,optional"`
	Size          int64  `json:"size,optional"`
	Url           string `json:"url,optional"`
	BucketName    string `json:"bucketName,optional"`
	IsLocalStored bool   `json:"isLocalStored,optional"`
	ContentType   string `json:"contentType,optional"`
	ResourceType  string `json:"resourceType,optional"`
}

type PivotDataDictionaryToObject struct {
	DataDictionaryType string `json:"dataDictionaryType,optional"`
	DataDictionaryKey  string `json:"dataDictionaryKey,optional"`
}

type ActivePriceEntry struct {
	Id        int64   `json:"id,optional"`
	UnitPrice float64 `json:"unitPrice,optional"`
	ListPrice float64 `json:"listPrice,optional"`
	Discount  float32 `json:"discount,optional"`
}

type ProductAttribute struct {
	Id         int64   `json:"id,optional"`
	Inventory  int16   `json:"inventory,optional"`
	SoldAmount int16   `json:"soldAmount,optional"`
	Weight     float32 `json:"weight,optional"`
	Volume     float32 `json:"volume,optional"`
	Encode     string  `json:"encode,optional"`
	BarCode    string  `json:"barCode,optional"`
	Extra      string  `json:"extra,optional"`
}

type Product struct {
	Id                     int64                          `json:"id,optional"`
	Name                   string                         `json:"name"`
	SPU                    string                         `json:"spu"`
	Type                   int                            `json:"type"`
	Plan                   int                            `json:"plan"`
	AccountingCategory     string                         `json:"accountingCategory"`
	CanSellOnline          bool                           `json:"canSellOnline,optional"`
	CanUseForDeduct        bool                           `json:"canUseForDeduct,optional"`
	Description            string                         `json:"description,optional"`
	AllowedSellQuantity    int                            `json:"purchasedQuantity,optional"`
	ValidityPeriodDays     int                            `json:"validityPeriodDays,optional"`
	SaleStartDate          string                         `json:"saleStartDate,optional"`
	SaleEndDate            string                         `json:"saleEndDate,optional"`
	ApprovalStatus         int                            `json:"approvalStatus,optional"`
	IsActivated            bool                           `json:"isActivated,optional,omitempty"`
	CreatedAt              string                         `json:"createdAt,optional"`
	ProductSpecifics       []*ProductSpecific             `json:"productSpecifics,optional"`
	PivotSalesChannels     []*PivotDataDictionaryToObject `json:"pivotSalesChannels,optional"`
	PivotPromoteChannels   []*PivotDataDictionaryToObject `json:"pivotPromoteChannels,optional"`
	ProductCategories      []*ProductCategory             `json:"productCategories,optional"`
	SalesChannelsItemIds   []int64                        `json:"salesChannelsItemIds,optional"`
	PromoteChannelsItemIds []int64                        `json:"promoteChannelsItemIds,optional"`
	CategoryIds            []int64                        `json:"categoryIds,optional"`
	CoverImageIds          []int64                        `json:"coverImageIds,optional"`
	CoverImages            []*ProductImage                `json:"coverImages,optional"`
	DetailImageIds         []int64                        `json:"detailImageIds,optional"`
	DetailImages           []*ProductImage                `json:"detailImages,optional"`
	ActivePriceEntry       *ActivePriceEntry              `json:"activePriceBookEntry,optional"`
	PriceBookEntries       []*PriceBookEntry              `json:"priceBookEntries,optional"`
	SKUs                   []*SKU                         `json:"skus,optional"`
	*ProductAttribute
	ViewedCount int `json:"viewedCount,optional"`
}

type ListProductsPageRequest struct {
	LikeName          string   `form:"likeName,optional"`
	ProductType       string   `form:"productType,optional"`
	Keys              []string `form:"keys,optional"`
	ProductCategoryId int      `form:"productCategoryId,optional"`
	OrderBy           string   `form:"orderBy,optional"`
	PageIndex         int      `form:"pageIndex,optional"`
	PageSize          int      `form:"pageSize,optional"`
}

type ListProductsPageReply struct {
	List      []Product `json:"list"`
	PageIndex int       `json:"pageIndex"`
	PageSize  int       `json:"pageSize"`
	Total     int64     `json:"total"`
}

type CreateProductRequest struct {
	Product
}

type CreateProductReply struct {
	ProductKey int64 `json:"id"`
}

type GetProductRequest struct {
	ProductId int64 `path:"id"`
}

type GetProductReply struct {
	*Product
}

type PutProductRequest struct {
	ProductId int64 `path:"id"`
	Product
}

type PutProductReply struct {
	*Product
}

type PatchProductRequest struct {
	ProductId int64 `path:"id"`
	Product
}

type PatchProductReply struct {
	*Product
}

type DeleteProductRequest struct {
	ProductId int64 `path:"id"`
}

type DeleteProductReply struct {
	ProductId int64 `json:"id"`
}

type AssignProductToProductCategoryRequest struct {
	Id                int64 `json:"id"`
	ProductCategoryId int64 `json:"productCategoryId"`
}

type AssignProductToProductCategoryReply struct {
	Product
}

type ImageAbleInfo struct {
	Icon            string `json:"icon"`
	BackgroundColor string `json:"backgroundColor"`
}

type CategoryImage struct {
	Id              int64  `json:"id,optional"`
	Filename        string `json:"filename,optional"`
	Size            int64  `json:"size,optional"`
	Url             string `json:"url,optional"`
	BucketName      string `json:"bucketName,optional"`
	IsLocalArtisand bool   `json:"isLocalArtisand,optional"`
	ContentType     string `json:"contentType,optional"`
	ResourceType    string `json:"resourceType,optional"`
}

type ProductCategory struct {
	Id          int64  `json:"id,optional"`
	PId         int64  `json:"pId"`
	Name        string `json:"name"`
	Sort        int    `json:"sort"`
	ViceName    string `json:"viceName"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt,optional"`
	ImageAbleInfo
	CoverImageId int64              `json:"coverImageId,optional"`
	CoverImage   *CategoryImage     `json:"coverImage,optional"`
	Children     []*ProductCategory `json:"children,optional"`
}

type ListProductCategoryTreeRequest struct {
	CategoryPId  int      `form:"categoryPId,optional"`
	NeedChildren bool     `form:"needChildren,optional"`
	Names        []string `form:"name,optional"`
	OrderBy      string   `form:"orderBy,optional"`
}

type ListProductCategoryTreeReply struct {
	ProductCategories []*ProductCategory `json:"tree"`
}

type CreateProductCategoryRequest struct {
	ProductCategory
}

type CreateProductCategoryReply struct {
	*ProductCategory
}

type UpdateProductCategoryRequest struct {
	Id int64 `path:"id"`
	ProductCategory
}

type UpdateProductCategoryReply struct {
	Id int64 `json:"id"`
}

type PatchProductCategoryRequest struct {
	Id  int64 `path:"id"`
	PId int64 `json:"pId"`
}

type PatchProductCategoryReply struct {
	ProductCategory
}

type GetProductCategoryRequest struct {
	ProductCategoryId int64 `path:"id"`
}

type GetProductCategoryReply struct {
	*ProductCategory
}

type DeleteProductCategoryRequest struct {
	Id int64 `path:"id"`
}

type DeleteProductCategoryReply struct {
	Id int64 `json:"id"`
}

type ProductSpecific struct {
	Id              int64             `json:"id,optional"`
	ProductId       int64             `json:"productId"`
	Name            string            `json:"name"`
	SpecificOptions []*SpecificOption `json:"specificOptions,optional"`
}

type SpecificOption struct {
	Id                int64  `json:"id,optional"`
	ProductSpecificId int64  `json:"ProductSpecificId,optional"`
	Name              string `json:"name,optional"`
	IsActivated       bool   `json:"isActivated,optional"`
}

type ListProductSpecificPageRequest struct {
	LikeName  string `form:"likeName,optional"`
	ProductId int64  `form:"productId"`
	OrderBy   string `form:"orderBy,optional"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type ListProductSpecificPageReply struct {
	List      []*ProductSpecific `json:"list"`
	PageIndex int                `json:"pageIndex"`
	PageSize  int                `json:"pageSize"`
	Total     int64              `json:"total"`
}

type CreateProductSpecificRequest struct {
	ProductSpecific
}

type CreateProductSpecificReply struct {
	ProductSpecificId int64 `json:"id"`
}

type ConfigProductSpecificRequest struct {
	ProductSpecifics []ProductSpecific `json:"productSpecifics"`
}

type ConfigProductSpecificReply struct {
	Result bool `json:"result"`
}

type GetProductSpecificRequest struct {
	ProductSpecificId int64 `path:"id"`
}

type GetProductSpecificReply struct {
	*ProductSpecific
}

type PutProductSpecificRequest struct {
	ProductSpecificId int64 `path:"id"`
	ProductSpecific
}

type PutProductSpecificReply struct {
	*ProductSpecific
}

type PatchProductSpecificRequest struct {
	ProductSpecificId int64 `path:"id"`
	ProductSpecific
}

type PatchProductSpecificReply struct {
	*ProductSpecific
}

type DeleteProductSpecificRequest struct {
	ProductSpecificId int64 `path:"id"`
}

type DeleteProductSpecificReply struct {
	ProductSpecificId int64 `json:"id"`
}

type SKU struct {
	Id         int64   `json:"id,optional"`
	UniqueId   string  `json:"uniqueId,optional"`
	SkuNo      string  `json:"skuNo,optional"`
	ProductId  int64   `json:"productId,optional"`
	Inventory  int     `json:"inventory,optional"`
	UnitPrice  float64 `json:"unitPrice,optional"`
	ListPrice  float64 `json:"listPrice,optional"`
	IsActive   bool    `json:"isActive,optional"`
	OptionsIds []int64 `json:"optionsIds,optional"`
}

type ListSKUPageRequest struct {
	LikeName  string `form:"likeName,optional"`
	ProductId int64  `form:"productId"`
	OrderBy   string `form:"orderBy,optional"`
	PageIndex int    `form:"pageIndex,optional"`
	PageSize  int    `form:"pageSize,optional"`
}

type ListSKUPageReply struct {
	List      []*SKU `json:"list"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
	Total     int64  `json:"total"`
}

type CreateSKURequest struct {
	SKU
}

type CreateSKUReply struct {
	SKUId int64 `json:"id"`
}

type ConfigSKURequest struct {
	SKUs []SKU `json:"skus"`
}

type ConfigSKUReply struct {
	Result bool `json:"result"`
}

type GetSKURequest struct {
	SKUId int64 `path:"id"`
}

type GetSKUReply struct {
	*SKU
}

type PutSKURequest struct {
	SKUId int64 `path:"id"`
	SKU
}

type PutSKUReply struct {
	*SKU
}

type PatchSKURequest struct {
	SKUId int64 `path:"id"`
	SKU
}

type PatchSKUReply struct {
	*SKU
}

type DeleteSKURequest struct {
	SKUId int64 `path:"id"`
}

type DeleteSKUReply struct {
	SKUId int64 `json:"id"`
}

type PriceConfig struct {
	Discount         float32 `json:"discount,optional"`
	Price            float64 `json:"price,optional"`
	Days             int8    `json:"days,optional"`
	Type             int8    `json:"type,optional"`
	PriceBookEntryId int64   `json:"priceBookEntryId,optional"`
	StartDate        string  `json:"startDate,optional"`
	EndDate          string  `json:"endDate,optional"`
}

type PriceBookEntrySpecific struct {
	Inventory int16   `json:"inventory,optional"`
	Weight    float32 `json:"weight,optional"`
	Volume    float32 `json:"volume,optional"`
	Encode    string  `json:"encode,optional"`
	BarCode   string  `json:"barCode,optional"`
	Extra     string  `json:"extra,optional"`
}

type PriceBookEntry struct {
	Id            int64             `json:"id,optional"`
	UniqueID      string            `json:"uniqueID,optional"`
	PriceBookId   int64             `json:"priceBookId"`
	ProductId     int64             `json:"productId"`
	SkuId         int64             `json:"skuId,optional"`
	UnitPrice     float64           `json:"unitPrice"`
	ListPrice     float64           `json:"listPrice,optional"`
	IsActive      bool              `json:"isActive, optional"`
	PriceConfigs  []*PriceConfig    `json:"priceConfigs, optional"`
	SKUEntries    []*PriceBookEntry `json:"skuEntries, optional"`
	PriceBookName string            `json:"priceBookName,optional"`
	ProductName   string            `json:"productName,optional"`
	SPU           string            `json:"spu,optional"`
	Discount      float32           `json:"discount,optional"`
}

type ListPriceBookEntriesPageRequest struct {
	LikeName    string `form:"likeName,optional"`
	PriceBookId int64  `form:"priceBookId,optional"`
	ProductId   int64  `form:"productId,optional"`
	SkuId       int64  `form:"skuId,optional"`
	PageIndex   int    `form:"pageIndex,optional"`
	PageSize    int    `form:"pageSize,optional"`
}

type ListPriceBookEntriesPageReply struct {
	List      []*PriceBookEntry `json:"list"`
	PageIndex int               `json:"pageIndex"`
	PageSize  int               `json:"pageSize"`
	Total     int64             `json:"total"`
}

type GetPriceBookEntryRequest struct {
	PriceBookEntry int64 `path:"id"`
}

type GetPriceBookEntryReply struct {
	*PriceBookEntry
}

type ConfigPriceBookEntryRequest struct {
	PriceBookEntries []PriceBookEntry `json:"priceBookEntries"`
}

type ConfigPriceBookEntryReply struct {
	PriceBookEntries []*PriceBookEntry `json:"list"`
}

type UpdatePriceBookEntryRequest struct {
	Id int64 `path:"id"`
	PriceBookEntry
}

type UpdatePriceBookEntryReply struct {
	Id int64 `json:"id"`
}

type DeletePriceBookEntryRequest struct {
	Id int64 `path:"id"`
}

type DeletePriceBookEntryReply struct {
	Id int64 `json:"id"`
}

type ArtisanImage struct {
	Id              int64  `json:"id,optional"`
	Filename        string `json:"filename,optional"`
	Size            int64  `json:"size,optional"`
	Url             string `json:"url,optional"`
	BucketName      string `json:"bucketName,optional"`
	IsLocalArtisand bool   `json:"isLocalArtisand,optional"`
	ContentType     string `json:"contentType,optional"`
	ResourceType    string `json:"resourceType,optional"`
}

type Artisan struct {
	Id             int64           `json:"id,optional"`
	EmployeeId     int64           `json:"employeeId,optional"`
	Name           string          `json:"name,optional"`
	Level          int8            `json:"level"`
	Gender         bool            `json:"gender,optional"`
	Birthday       string          `json:"birthday,optional"`
	PhoneNumber    string          `json:"phoneNumber,optional"`
	WorkNo         string          `json:"workNo"`
	Email          string          `json:"email,optional"`
	Experience     string          `json:"experience,optional"`
	Specialty      string          `json:"specialty,optional"`
	Certificate    string          `json:"certificate,optional"`
	Address        string          `json:"address,optional"`
	CreatedAt      string          `json:"createdAt,optional"`
	CoverImageId   int64           `json:"coverImageId,optional"`
	CoverImage     *ArtisanImage   `json:"coverImage,optional"`
	DetailImageIds []int64         `json:"detailImageIds,optional"`
	DetailImages   []*ArtisanImage `json:"detailIOmages,optional"`
	StoreIds       []int64         `json:"storeIds,optional"`
}

type ListArtisansPageRequest struct {
	Ids       []int64 `form:"ids,optional"`
	StoreIds  []int64 `form:"storeIds,optional"`
	LikeName  string  `form:"likeName,optional"`
	OrderBy   string  `form:"orderBy,optional"`
	PageIndex int     `form:"pageIndex,optional"`
	PageSize  int     `form:"pageSize,optional"`
}

type ListArtisansPageReply struct {
	List      []*Artisan `json:"list"`
	PageIndex int        `json:"pageIndex"`
	PageSize  int        `json:"pageSize"`
	Total     int64      `json:"total"`
}

type CreateArtisanRequest struct {
	Artisan
}

type CreateArtisanReply struct {
	ArtisanId int64 `json:"id"`
}

type GetArtisanRequest struct {
	ArtisanId int64 `path:"id"`
}

type GetArtisanReply struct {
	*Artisan
}

type PutArtisanRequest struct {
	ArtisanId int64 `path:"id"`
	Artisan
}

type PutArtisanReply struct {
	*Artisan
}

type DeleteArtisanRequest struct {
	ArtisanId int64 `path:"id"`
}

type DeleteArtisanReply struct {
	ArtisanId int64 `json:"id"`
}

type ShippingAddress struct {
	Id           int64  `json:"id,optional"`
	CustomerId   int64  `json:"customerId,optional"`
	Name         string `json:"name,optional"`
	Recipient    string `json:"recipient,optional"`
	AddressLine  string `json:"addressLine,optional"`
	AddressLine2 string `json:"addressLine2,optional"`
	Street       string `json:"street,optional"`
	City         string `json:"city,optional"`
	Province     string `json:"province,optional"`
	PostalCode   string `json:"postalCode,optional"`
	Country      string `json:"country,optional"`
	PhoneNumber  string `json:"phoneNumber,optional"`
	IsDefault    bool   `json:"isDefault,optional"`
}

type ListShippingAddressesPageRequest struct {
	NameLike  []string `form:"nameLike,optional"`
	OrderBy   string   `form:"orderBy,optional"`
	PageIndex int      `form:"pageIndex,optional"`
	PageSize  int      `form:"pageSize,optional"`
}

type ListShippingAddressesPageReply struct {
	List      []*ShippingAddress `json:"list"`
	PageIndex int                `json:"pageIndex"`
	PageSize  int                `json:"pageSize"`
	Total     int64              `json:"total"`
}

type CreateShippingAddressRequest struct {
	ShippingAddress
}

type CreateShippingAddressReply struct {
	*ShippingAddress
}

type GetShippingAddressRequest struct {
	ShippingAddressId int64 `path:"id"`
}

type GetShippingAddressReply struct {
	*ShippingAddress
}

type PutShippingAddressRequest struct {
	ShippingAddressId int64 `path:"id"`
	ShippingAddress
}

type PutShippingAddressReply struct {
	*ShippingAddress
}

type PatchShippingAddressRequest struct {
	ShippingAddressId int64 `path:"id"`
	ShippingAddress
}

type PatchShippingAddressReply struct {
	*ShippingAddress
}

type DeleteShippingAddressRequest struct {
	ShippingAddressId int64 `path:"id"`
}

type DeleteShippingAddressReply struct {
	ShippingAddressId int64 `json:"id"`
}

type BillingAddress struct {
	Id           int64  `json:"id,optional"`
	OrderId      int64  `gorm:"comment:订单Id; index" json:"orderId"`
	CustomerId   int64  `json:"customerId,optional"`
	Recipient    string `json:"recipient,optional"`
	AddressLine  string `json:"addressLine,optional"`
	AddressLine2 string `json:"addressLine2,optional"`
	Street       string `json:"street,optional"`
	City         string `json:"city,optional"`
	Province     string `json:"province,optional"`
	PostalCode   string `json:"postalCode,optional"`
	Country      string `json:"country,optional"`
	PhoneNumber  string `json:"phoneNumber,optional"`
	IsDefault    bool   `json:"isDefault,optional"`
}

type ListBillingAddressesPageRequest struct {
	NameLike  []string `form:"nameLike,optional"`
	OrderBy   string   `form:"orderBy,optional"`
	PageIndex int      `form:"pageIndex,optional"`
	PageSize  int      `form:"pageSize,optional"`
}

type ListBillingAddressesPageReply struct {
	List      []*BillingAddress `json:"list"`
	PageIndex int               `json:"pageIndex"`
	PageSize  int               `json:"pageSize"`
	Total     int64             `json:"total"`
}

type CreateBillingAddressRequest struct {
	BillingAddress *BillingAddress `json:"billingAddress"`
}

type CreateBillingAddressReply struct {
	BillingAddressId int64 `json:"id"`
}

type GetBillingAddressRequest struct {
	BillingAddressId int64 `path:"id"`
}

type GetBillingAddressReply struct {
	BillingAddress *BillingAddress `json:"billingAddress"`
}

type PutBillingAddressRequest struct {
	BillingAddressId int64           `path:"id"`
	BillingAddress   *BillingAddress `json:"billingAddress"`
}

type PutBillingAddressReply struct {
	BillingAddress *BillingAddress `json:"billingAddress"`
}

type PatchBillingAddressRequest struct {
	BillingAddressId int64           `path:"id"`
	BillingAddress   *BillingAddress `json:"billingAddress"`
}

type PatchBillingAddressReply struct {
	BillingAddress *BillingAddress `json:"billingAddress"`
}

type DeleteBillingAddressRequest struct {
	BillingAddressId int64 `path:"id"`
}

type DeleteBillingAddressReply struct {
	BillingAddressId int64 `json:"id"`
}

type DeliveryAddress struct {
	Id           int64  `json:"id,optional"`
	OrderId      int64  `gorm:"comment:订单Id; index" json:"orderId"`
	CustomerId   int64  `json:"customerId,optional"`
	Recipient    string `json:"recipient,optional"`
	AddressLine  string `json:"addressLine,optional"`
	AddressLine2 string `json:"addressLine2,optional"`
	Street       string `json:"street,optional"`
	City         string `json:"city,optional"`
	Province     string `json:"province,optional"`
	PostalCode   string `json:"postalCode,optional"`
	Country      string `json:"country,optional"`
	PhoneNumber  string `json:"phoneNumber,optional"`
	IsDefault    bool   `json:"isDefault,optional"`
}

type ListDeliveryAddressesPageRequest struct {
	NameLike  []string `form:"nameLike,optional"`
	OrderBy   string   `form:"orderBy,optional"`
	PageIndex int      `form:"pageIndex,optional"`
	PageSize  int      `form:"pageSize,optional"`
}

type ListDeliveryAddressesPageReply struct {
	List      []*DeliveryAddress `json:"list"`
	PageIndex int                `json:"pageIndex"`
	PageSize  int                `json:"pageSize"`
	Total     int64              `json:"total"`
}

type CreateDeliveryAddressRequest struct {
	DeliveryAddress *DeliveryAddress `json:"deliveryAddress"`
}

type CreateDeliveryAddressReply struct {
	DeliveryAddressId int64 `json:"id"`
}

type GetDeliveryAddressRequest struct {
	DeliveryAddressId int64 `path:"id"`
}

type GetDeliveryAddressReply struct {
	DeliveryAddress *DeliveryAddress `json:"deliveryAddress"`
}

type PutDeliveryAddressRequest struct {
	DeliveryAddressId int64            `path:"id"`
	DeliveryAddress   *DeliveryAddress `json:"deliveryAddress"`
}

type PutDeliveryAddressReply struct {
	DeliveryAddress *DeliveryAddress `json:"deliveryAddress"`
}

type PatchDeliveryAddressRequest struct {
	DeliveryAddressId int64            `path:"id"`
	DeliveryAddress   *DeliveryAddress `json:"deliveryAddress"`
}

type PatchDeliveryAddressReply struct {
	DeliveryAddress *DeliveryAddress `json:"deliveryAddress"`
}

type DeleteDeliveryAddressRequest struct {
	DeliveryAddressId int64 `path:"id"`
}

type DeleteDeliveryAddressReply struct {
	DeliveryAddressId int64 `json:"id"`
}

type Warehouse struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Region        string `json:"region"`
	Type          string `json:"type"`
	Capacity      int64  `json:"capacity"`
	ContactPerson string `json:"contactPerson"`
	ContactPhone  string `json:"contactPhone"`
	IsActive      bool   `json:"isActive, omitempty"`
}

type ListWarehousesRequest struct {
	NameLike  string `json:"nameLike"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type ListWarehousesResponse struct {
	List      []*Warehouse `json:"list"`
	PageIndex int          `json:"pageIndex"`
	PageSize  int          `json:"pageSize"`
	Total     int64        `json:"total"`
}

type GetWarehouseRequest struct {
	Id int64 `json:"id"`
}

type GetWarehouseResponse struct {
	Warehouse *Warehouse `json:"warehouse"`
}

type CreateWarehouseRequest struct {
	Warehouse *Warehouse `json:"warehouse"`
}

type CreateWarehouseResponse struct {
	WarehouseId int64 `json:"warehouseId"`
}

type UpdateWarehouseRequest struct {
	Id        int64      `json:"id"`
	Warehouse *Warehouse `json:"warehouse"`
}

type UpdateWarehouseResponse struct {
	Warehouse *Warehouse `json:"warehouse"`
}

type PatchWarehouseRequest struct {
	Id        int64      `json:"id"`
	Warehouse *Warehouse `json:"warehouse"`
}

type PatchWarehouseResponse struct {
	Warehouse *Warehouse `json:"warehouse"`
}

type DeleteWarehouseRequest struct {
	Id int64 `json:"id"`
}

type DeleteWarehouseResponse struct {
	WarehouseId int64 `json:"warehouseId"`
}

type ContractWayGroupNode struct {
	Id        int64                  `json:"id"`
	GroupName string                 `json:"groupName"`
	Children  []ContractWayGroupNode `json:"children"`
}

type GetContractWayGroupTreeRequest struct {
}

type GetContractWayGroupTreeReply struct {
	GroupTree ContractWayGroupNode `json:"tree"`
}

type ContractWayGroup struct {
	Id        int64  `json:"id"`
	GroupName string `json:"groupName"`
}

type GetContractWayGroupListRequest struct {
	GroupName string `form:"groupName,optional"`
}

type GetContractWayGroupListReply struct {
	Groups []ContractWayGroup `json:"groups"`
}

type GetContractWaysRequest struct {
	EmployeeId int64  `form:"employeeId,optional"`
	Name       string `form:"name,optional"`
	StartDate  string `form:"startDate,optional"`
	EndDate    string `form:"endDate,optional"`
	PageIndex  int    `form:"pageIndex"`
	PageSize   int    `form:"pageSize"`
}

type GetContractWaysReply struct {
	List      []ContractWay `json:"list"`
	PageIndex int           `json:"pageIndex"`
	PageSize  int           `json:"pageSize"`
	Total     int64         `json:"total"`
}

type ContractWay struct {
	Id            int64    `json:"id"`
	Type          int      `json:"type"`
	Scene         int      `json:"scene"`
	Style         string   `json:"style,optional"`
	Remark        string   `json:"remark,optional"`
	SkipVerify    bool     `json:"skipVerify,optional"`
	State         string   `json:"state,optional"`
	Users         []string `json:"users,optional"`
	Parties       []int64  `json:"parties,optional"`
	IsTemp        bool     `json:"isTemp,optional"`
	ExpiresIn     int      `json:"expiresIn,optional"`
	ChatExpiresIn int      `json:"chatExpiresIn,optional"`
	UnionId       string   `json:"unionId,optional"`
	IsExclusive   bool     `json:"isExclusive,optional"`
	Conclusions   string   `json:"conclusions,optional"`
}

type CreateContractWayRequest struct {
	Type          int      `json:"type"`
	Scene         int      `json:"scene"`
	Style         string   `json:"style,optional"`
	Remark        string   `json:"remark,optional"`
	SkipVerify    bool     `json:"skipVerify,optional"`
	State         string   `json:"state,optional"`
	Users         []string `json:"users,optional"`
	Parties       []int64  `json:"parties,optional"`
	IsTemp        bool     `json:"isTemp,optional"`
	ExpiresIn     int      `json:"expiresIn,optional"`
	ChatExpiresIn int      `json:"chatExpiresIn,optional"`
	UnionId       string   `json:"unionId,optional"`
	IsExclusive   bool     `json:"isExclusive,optional"`
	Conclusions   string   `json:"conclusions,optional"`
}

type CreateContractWayReply struct {
	Id int64 `json:"id"`
}

type UpdateContractWayRequest struct {
	Id            int64    `path:"id"`
	Type          int      `json:"type,optional"`
	Scene         int      `json:"scene,optional"`
	Style         string   `json:"style,optional"`
	Remark        string   `json:"remark,optional"`
	SkipVerify    bool     `json:"skipVerify,optional"`
	State         string   `json:"state,optional"`
	Users         []string `json:"users,optional"`
	Parties       []int64  `json:"parties,optional"`
	IsTemp        bool     `json:"isTemp,optional"`
	ExpiresIn     int      `json:"expiresIn,optional"`
	ChatExpiresIn int      `json:"chatExpiresIn,optional"`
	UnionId       string   `json:"unionId,optional"`
	IsExclusive   bool     `json:"isExclusive,optional"`
	Conclusions   string   `json:"conclusions,optional"`
}

type UpdateContractWayReply struct {
	ContractWayUpdated ContractWay `json:"contractWayUpdated"`
}

type DeleteContractWayRequest struct {
	Id int64 `path:"id"`
}

type DeleteContractWayReply struct {
	Id int64 `json:"id"`
}

type WeWorkCustomer struct {
	Name            string   `json:"name"`
	AdderId         int64    `json:"adderId"`
	AddTime         string   `json:"addTime"`
	PatchTime       string   `json:"updateTime"`
	AddChannel      string   `json:"addChannel"`
	TagGroupIdList  []int64  `json:"tagGroupIdList"`
	TagIdList       []int64  `json:"tagIdList"`
	PersonalTagList []string `json:"personalTagList"`
	Age             int      `json:"age"`
	Email           string   `json:"email"`
	PhoneNumber     string   `json:"phoneNumber"`
	Address         string   `json:"address"`
	Birthday        string   `json:"birthday"`
	Remark          string   `json:"remark"`
	GroupChatId     int64    `json:"groupChatId"`
}

type GetWeWorkCustomerRequest struct {
	Id string `path:"id"`
}

type GetWeWorkCustomerReply struct {
	WeWorkCustomer
}

type ListWeWorkCustomersRequest struct {
	LikeName     string `form:"likeName,optional"`
	FollowUserId string `form:"followUserId,optional"`
	AddWay       int    `form:"addWay,optional"`
	PageIndex    int    `form:"pageIndex,optional"`
	PageSize     int    `form:"pageSize,optional"`
}

type ListWeWorkCustomersReply struct {
	List      []GetWeWorkCustomerReply `json:"list"`
	PageIndex int                      `json:"pageIndex"`
	PageSize  int                      `json:"pageSize"`
	Total     int64                    `json:"total"`
}

type PatchWeWorkCustomerRequest struct {
	Id          string `path:"id"`
	Name        string `json:"name,optional"`
	Age         int    `json:"age,optional"`
	Email       string `json:"email,optional"`
	PhoneNumber string `json:"phoneNumber,optional"`
	Address     string `json:"address,optional"`
	Birthday    string `json:"birthday,optional"`
	Remark      string `json:"remark,optional"`
	GroupChatId int64  `json:"groupChatId,optional"`
}

type PatchWeWorkCustomerReply struct {
	WeWorkCustomer
}

type SyncWeWorkCustomerReply struct {
	Status string `json:"status"`
}

type SyncWeWorkContactReply struct {
	Status string `json:"status"`
}

type ListWeWorkEmployeeReqeust struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}

type ListWeWorkEmployeeReply struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	Total     int `json:"total"`
}

type MPCustomerLoginRequest struct {
	Code string `json:"code"`
}

type MPCustomerAuthRequest struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type MPCustomerLoginAuthReply struct {
	OpenId      string  `json:"openId"`
	UnionId     string  `json:"unionId"`
	PhoneNumber string  `json:"phoneNumber"`
	NickName    string  `json:"nickName"`
	AvatarURL   string  `json:"avatarURL"`
	Gender      string  `json:"gender"`
	Token       MPToken `json:"token"`
}

type MPToken struct {
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type ListProductCategoriesRequest struct {
	CategoryPId  int  `form:"categoryPId,optional"`
	NeedChildren bool `form:"needChildren,optional"`
	Limit        int  `form:"limit,optional"`
}

type ListProductCategoriesReply struct {
	ProductCategories []*ProductCategory `json:"list"`
}

type Cart struct {
	Id         int64       `json:"id", optional"`
	CustomerId int64       `json:"customerId", optional"`
	Status     int         `json:"status", optional"`
	Items      []*CartItem `json:"items", optional"`
}

type CartItem struct {
	Id             int64   `json:"id, optional"`
	CustomerId     int64   `json:"customerId,omitempty,optional"`
	CartId         int64   `json:"cartId,omitempty,optional"`
	ProductId      int64   `json:"productId,omitempty,optional"`
	SkuId          int64   `json:"skuId,omitempty,optional"`
	ProductName    string  `json:"productName,omitempty,optional"`
	ListPrice      float64 `json:"listPrice,omitempty,optional"`
	UnitPrice      float64 `json:"unitPrice,omitempty,optional"`
	Discount       float64 `json:"discount,omitempty,optional"`
	Quantity       int     `json:"quantity,omitempty,optional"`
	Specifications string  `json:"specifications,omitempty,optional"`
	ImageURL       string  `json:"imageUrl,omitempty,optional"`
}

type ListCartItemsPageRequest struct {
	PageIndex int `form:"pageIndex,optional"`
	PageSize  int `form:"pageSize,optional"`
}

type ListCartItemsPageReply struct {
	List      []*CartItem `json:"list"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Total     int64       `json:"total"`
}

type GetCartRequest struct {
}

type GetCartReply struct {
	*Cart
}

type AddToCartRequest struct {
	CartItem
}

type AddToCartReply struct {
	*CartItem
}

type UpdateCartItemQuantityRequest struct {
	ItemId   int64 `path:"itemId"`
	Quantity int   `json:"quantity"`
}

type UpdateCartItemQuantityReply struct {
	*CartItem
}

type RemoveCartItemRequest struct {
	ItemId int64 `path:"itemId"`
}

type RemoveCartItemReply struct {
	ItemId int64 `json:"itemId"`
}

type ClearCartItemsRequest struct {
}

type ClearCartItemsReply struct {
	Message string `json:"message"`
}

type CreateOrderByProductsRequest struct {
	PriceBookId       int64   `json:"PriceBookId,optional,emptyomit"`
	ProductIds        []int64 `json:"productIds"`
	SkuIds            []int64 `json:"skuIds"`
	Quantities        []int   `json:"quantities"`
	ShippingAddressId int64   `json:"shippingAddressId,optional,emptyomit"`
	Comment           string  `json:"comment"`
}

type CreateOrderByProductsReply struct {
	OrderId       int64   `json:"orderId"`
	PaymentAmount float64 `json:"paymentAmount"`
}

type CreateOrderByCartItemsRequest struct {
	CartItemIds       []int64 `json:"cartItemIds"`
	ShippingAddressId int64   `json:"shippingAddressId"`
	Comment           string  `json:"comment"`
}

type CreateOrderByCartItemsReply struct {
	OrderId       int64   `json:"orderId"`
	CartId        int64   `json:"cartId"`
	PaymentAmount float64 `json:"paymentAmount"`
}

type CancelOrderRequest struct {
	OrderId int64 `path:"id"`
}

type CancelOrderReply struct {
	OrderId int64 `json:"orderId,optional"`
}

type OrderItem struct {
	Id               int64   `json:"id,optional"`
	OrderId          int64   `json:"orderId,optional"`
	PriceBookEntryId int64   `json:"priceBookEntryId,optional"`
	CustomerId       int64   `json:"customerId,optional"`
	Type             int     `json:"type,optional"`
	Status           int     `json:"status,optional"`
	Quantity         int     `json:"quantity,optional"`
	UnitPrice        float64 `json:"unitPrice,optional"`
	ListPrice        float64 `json:"listPrice,optional"`
	SellingPrice     float64 `json:"sellingPrice,optional"`
	CoverUrl         string  `json:"coverUrl,optional"`
	ProdcutName      string  `json:"productName,optional"`
	SkuNo            string  `json:"skuNo,optional"`
}

type Order struct {
	Id             int64        `json:"id,optional"`
	CustomerId     int64        `json:"customerId,optional"`
	CartId         int64        `json:"cartId,optional"`
	PaymentType    int          `json:"paymentType,optional"`
	Type           int          `json:"type,optional"`
	Status         int          `json:"status,optional"`
	OrderNumber    string       `json:"orderNumber,optional"`
	Discount       float64      `json:"discount,optional"`
	ListPrice      float64      `json:"listPrice,optional"`
	UnitPrice      float64      `json:"unitPrice,optional"`
	Comment        string       `json:"comment,optional"`
	CompletedAt    string       `json:"completedAt,optional,omitempty"`
	CancelledAt    string       `json:"cancelledAt,optional,omitempty"`
	ShippingMethod string       `json:"shippingMethod,optional,omitempty"`
	CreatedAt      string       `json:"createdAt,optional,omitempty"`
	OrderItems     []*OrderItem `json:"orderItems,optional"`
	Payments       []*Payment   `json:"payments,optional"`
}

type ListOrdersPageRequest struct {
	OrderType   []int    `form:"orderType,optional,omitempty"`
	OrderStatus []int    `form:"orderStatus,optional,omitempty"`
	Keys        []string `form:"keys,optional,omitempty"`
	OrderBy     string   `form:"orderBy,optional,omitempty"`
	PageIndex   int      `form:"pageIndex,optional,omitempty"`
	PageSize    int      `form:"pageSize,optional,omitempty"`
}

type ListOrdersPageReply struct {
	List      []*Order `json:"list"`
	PageIndex int      `json:"pageIndex"`
	PageSize  int      `json:"pageSize"`
	Total     int64    `json:"total"`
}

type ExportOrdersRequest struct {
	OrderType   []int  `form:"orderType,optional,omitempty"`
	OrderStatus []int  `form:"orderStatus,optional,omitempty"`
	OrderBy     string `form:"orderBy,optional"`
	StartAt     string `form:"startAt"`
	EndAt       string `form:"endAt"`
}

type ExportOrdersReply struct {
	Content  []byte `json:"content"`
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
	FileType string `json:"fileType"`
}

type CreateOrderRequest struct {
	Order
}

type CreateOrderReply struct {
	OrderId int64 `json:"orderId,omitempty"`
}

type GetOrderRequest struct {
	OrderId int64 `path:"id"`
}

type GetOrderReply struct {
	*Order
}

type PutOrderRequest struct {
	OrderId int64 `path:"id"`
	Order
}

type PutOrderReply struct {
	*Order
}

type PatchOrderRequest struct {
	OrderId int64 `path:"id"`
	Order
}

type PatchOrderReply struct {
	*Order
}

type DeleteOrderRequest struct {
	OrderId int64 `path:"id"`
}

type DeleteOrderReply struct {
	OrderId int64 `path:"id"`
}

type PaymentItem struct {
	Id                  int64   `json:"id,optional"`
	PaymentID           int64   `json:"paymentID,optional"`
	Quantity            int     `json:"quantity,optional"`
	UnitPrice           float64 `json:"unitPrice,optional"`
	PaymentCustomerName string  `json:"paymentCustomerName,optional"`
	BankInformation     string  `json:"bankInformation,optional"`
	BankResponseCode    string  `json:"bankResponseCode,optional"`
	CarrierType         string  `json:"carrierType,optional"`
	CreditCardNumber    string  `json:"creditCardNumber,optional"`
	DeductMembershipId  string  `json:"deductMembershipId,optional"`
	DeductionPoint      int32   `json:"deductionPoint,optional"`
	InvoiceCreateTime   string  `json:"invoiceCreateTime,optional"`
	InvoiceNumber       string  `json:"invoiceNumber,optional"`
	InvoiceTotalAmount  float64 `json:"invoiceTotalAmount,optional"`
	TaxIdNumber         string  `json:"taxIdNumber,optional"`
}

type Payment struct {
	Id              int64          `json:"id,optional"`
	OrderId         int64          `json:"orderId,optional"`
	PaymentDate     string         `json:"paymentDate,optional"`
	PaymentType     int            `json:"paymentType,optional"`
	PaidAmount      float64        `json:"paidAmount,optional"`
	PaymentNumber   string         `json:"paymentNumber,optional"`
	ReferenceNumber string         `json:"referenceNumber,optional"`
	Status          int            `json:"status,optional"`
	PaymentItems    []*PaymentItem `json:"paymentItems,optional"`
}

type ListPaymentsPageRequest struct {
	PaymentType string   `form:"paymentType,optional"`
	Keys        []string `form:"keys,optional"`
	OrderBy     string   `form:"orderBy,optional"`
	PageIndex   int      `form:"pageIndex,optional"`
	PageSize    int      `form:"pageSize,optional"`
}

type ListPaymentsPageReply struct {
	List      []*Payment `json:"list"`
	PageIndex int        `json:"pageIndex"`
	PageSize  int        `json:"pageSize"`
	Total     int64      `json:"total"`
}

type CreatePaymentRequest struct {
	Payment
}

type CreatePaymentReply struct {
	PaymentKey int64 `json:"id"`
}

type GetPaymentRequest struct {
	PaymentId int64 `path:"id"`
}

type GetPaymentReply struct {
	*Payment
}

type PutPaymentRequest struct {
	PaymentId int64 `path:"id"`
	Payment
}

type PutPaymentReply struct {
	*Payment
}

type PatchPaymentRequest struct {
	PaymentId int64 `path:"id"`
	Payment
}

type PatchPaymentReply struct {
	*Payment
}

type DeletePaymentRequest struct {
	PaymentId int64 `path:"id"`
}

type DeletePaymentReply struct {
	PaymentId int64 `json:"id"`
}

type CreatePaymentFromOrderRequest struct {
	OrderId     int64  `json:"orderId"`
	PaymentType int    `json:"paymentType"`
	Comment     string `json:"comment,optional"`
}

type CreatePaymentFromOrderRequestReply struct {
	PaymentId int64       `json:"paymentId"`
	Data      interface{} `json:"data"`
}

type UpdatePaymentRequest struct {
	PaymentId int64 `path:"id"`
	Payment
}

type UpdatePaymentReply struct {
	*Payment
}

type CustomerLoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type CustomerRegisterRequest struct {
	Account    string `json:"account"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

type CustomerRegisterReply struct {
	CustomerId int64 `json:"customerId"`
}

type CustomerRegisterByPhoneRequest struct {
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

type CustomerRegisterByPhoneReply struct {
	CustomerId int64 `json:"customerId"`
}

type UpdateCustomerProfileRequest struct {
	CustomerId int64 `path:"id"`
	Customer
}

type UpdateCustomerProfileReply struct {
	*Customer
}

type CustomerLoginAuthReply struct {
	OpenId      string   `json:"openId"`
	UnionId     string   `json:"unionId"`
	PhoneNumber string   `json:"phoneNumber"`
	NickName    string   `json:"nickName"`
	AvatarURL   string   `json:"avatarURL"`
	Gender      string   `json:"gender"`
	Token       WebToken `json:"token"`
}

type WebToken struct {
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type OACustomerLoginRequest struct {
	Code string `json:"code"`
}

type OACustomerAuthRequest struct {
	Code          string `json:"code"`
	IV            string `json:"iv"`
	EncryptedData string `json:"encryptedData"`
}

type OACustomerLoginAuthReply struct {
	OpenId      string  `json:"openId"`
	UnionId     string  `json:"unionId"`
	PhoneNumber string  `json:"phoneNumber"`
	NickName    string  `json:"nickName"`
	AvatarURL   string  `json:"avatarURL"`
	Gender      string  `json:"gender"`
	Token       OAToken `json:"token"`
}

type OAToken struct {
	TokenType    string `json:"tokenType"`
	ExpiresIn    string `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Route struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type RegisterPluginRequest struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
	Addr   string  `json:"addr"`
}

type RegisterPluginReply struct {
	Name string                 `json:"name"`
	Etc  map[string]interface{} `json:"etc"`
}

type PluginWebRouteMeta struct {
	Locale      string `json:"locale"`
	Icon        string `json:"icon"`
	RequestAuth bool   `json:"requestAuth"`
}

type PluginWebRoutes struct {
	Name     string             `json:"name"`
	Path     string             `json:"path"`
	Meta     PluginWebRouteMeta `json:"meta"`
	Children []PluginWebRoutes  `json:"children"`
}

type ListPluginRequest struct {
	Socpe []string `form:"scope"`
}

type PluginWebInfo struct {
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Verison   string  `json:"version"`
	IsEnabled bool    `json:"isEnabled"`
	Routes    []Route `json:"routes"`
}

type ListPluginReply struct {
	Plugins []PluginWebInfo `json:"plugins"`
}

type ListPluginFrontendRoutesReply struct {
	Routes []PluginWebRoutes `json:"routes"`
}

type HealthCheckRequest struct {
}

type HealthCheckReply struct {
}
