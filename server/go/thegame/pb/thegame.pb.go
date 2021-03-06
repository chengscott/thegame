// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thegame.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ability int32

const (
	Ability_health_regen       Ability = 0
	Ability_max_health         Ability = 1
	Ability_body_damage        Ability = 2
	Ability_bullet_speed       Ability = 3
	Ability_bullet_penetration Ability = 4
	Ability_bullet_damage      Ability = 5
	Ability_bullet_reload      Ability = 6
	Ability_movement_speed     Ability = 7
)

var Ability_name = map[int32]string{
	0: "health_regen",
	1: "max_health",
	2: "body_damage",
	3: "bullet_speed",
	4: "bullet_penetration",
	5: "bullet_damage",
	6: "bullet_reload",
	7: "movement_speed",
}

var Ability_value = map[string]int32{
	"health_regen":       0,
	"max_health":         1,
	"body_damage":        2,
	"bullet_speed":       3,
	"bullet_penetration": 4,
	"bullet_damage":      5,
	"bullet_reload":      6,
	"movement_speed":     7,
}

func (x Ability) String() string {
	return proto.EnumName(Ability_name, int32(x))
}

func (Ability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{0}
}

type Controls struct {
	Accelerate            bool      `protobuf:"varint,1,opt,name=accelerate,proto3" json:"accelerate,omitempty"`
	AccelerationDirection float64   `protobuf:"fixed64,2,opt,name=acceleration_direction,json=accelerationDirection,proto3" json:"acceleration_direction,omitempty"`
	Shoot                 bool      `protobuf:"varint,3,opt,name=shoot,proto3" json:"shoot,omitempty"`
	ShootDirection        float64   `protobuf:"fixed64,4,opt,name=shoot_direction,json=shootDirection,proto3" json:"shoot_direction,omitempty"`
	LevelUp               []Ability `protobuf:"varint,5,rep,packed,name=level_up,json=levelUp,proto3,enum=Ability" json:"level_up,omitempty"`
	Name                  string    `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}  `json:"-"`
	XXX_unrecognized      []byte    `json:"-"`
	XXX_sizecache         int32     `json:"-"`
}

func (m *Controls) Reset()         { *m = Controls{} }
func (m *Controls) String() string { return proto.CompactTextString(m) }
func (*Controls) ProtoMessage()    {}
func (*Controls) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{0}
}

func (m *Controls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Controls.Unmarshal(m, b)
}
func (m *Controls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Controls.Marshal(b, m, deterministic)
}
func (m *Controls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Controls.Merge(m, src)
}
func (m *Controls) XXX_Size() int {
	return xxx_messageInfo_Controls.Size(m)
}
func (m *Controls) XXX_DiscardUnknown() {
	xxx_messageInfo_Controls.DiscardUnknown(m)
}

var xxx_messageInfo_Controls proto.InternalMessageInfo

func (m *Controls) GetAccelerate() bool {
	if m != nil {
		return m.Accelerate
	}
	return false
}

func (m *Controls) GetAccelerationDirection() float64 {
	if m != nil {
		return m.AccelerationDirection
	}
	return 0
}

func (m *Controls) GetShoot() bool {
	if m != nil {
		return m.Shoot
	}
	return false
}

func (m *Controls) GetShootDirection() float64 {
	if m != nil {
		return m.ShootDirection
	}
	return 0
}

func (m *Controls) GetLevelUp() []Ability {
	if m != nil {
		return m.LevelUp
	}
	return nil
}

func (m *Controls) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ViewRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ViewRequest) Reset()         { *m = ViewRequest{} }
func (m *ViewRequest) String() string { return proto.CompactTextString(m) }
func (*ViewRequest) ProtoMessage()    {}
func (*ViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{1}
}

func (m *ViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewRequest.Unmarshal(m, b)
}
func (m *ViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewRequest.Marshal(b, m, deterministic)
}
func (m *ViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewRequest.Merge(m, src)
}
func (m *ViewRequest) XXX_Size() int {
	return xxx_messageInfo_ViewRequest.Size(m)
}
func (m *ViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ViewRequest proto.InternalMessageInfo

func (m *ViewRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GameState struct {
	Meta                 *GameState_Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Polygons             []*Polygon      `protobuf:"bytes,2,rep,name=polygons,proto3" json:"polygons,omitempty"`
	Bullets              []*Bullet       `protobuf:"bytes,3,rep,name=bullets,proto3" json:"bullets,omitempty"`
	Heroes               []*Hero         `protobuf:"bytes,4,rep,name=heroes,proto3" json:"heroes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GameState) Reset()         { *m = GameState{} }
func (m *GameState) String() string { return proto.CompactTextString(m) }
func (*GameState) ProtoMessage()    {}
func (*GameState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{2}
}

func (m *GameState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameState.Unmarshal(m, b)
}
func (m *GameState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameState.Marshal(b, m, deterministic)
}
func (m *GameState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameState.Merge(m, src)
}
func (m *GameState) XXX_Size() int {
	return xxx_messageInfo_GameState.Size(m)
}
func (m *GameState) XXX_DiscardUnknown() {
	xxx_messageInfo_GameState.DiscardUnknown(m)
}

var xxx_messageInfo_GameState proto.InternalMessageInfo

func (m *GameState) GetMeta() *GameState_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GameState) GetPolygons() []*Polygon {
	if m != nil {
		return m.Polygons
	}
	return nil
}

func (m *GameState) GetBullets() []*Bullet {
	if m != nil {
		return m.Bullets
	}
	return nil
}

func (m *GameState) GetHeroes() []*Hero {
	if m != nil {
		return m.Heroes
	}
	return nil
}

type GameState_Meta struct {
	HeroId               int32         `protobuf:"varint,2,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	CenterPosition       *Vector       `protobuf:"bytes,5,opt,name=center_position,json=centerPosition,proto3" json:"center_position,omitempty"`
	Scores               []*ScoreEntry `protobuf:"bytes,4,rep,name=scores,proto3" json:"scores,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameState_Meta) Reset()         { *m = GameState_Meta{} }
func (m *GameState_Meta) String() string { return proto.CompactTextString(m) }
func (*GameState_Meta) ProtoMessage()    {}
func (*GameState_Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{2, 0}
}

func (m *GameState_Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameState_Meta.Unmarshal(m, b)
}
func (m *GameState_Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameState_Meta.Marshal(b, m, deterministic)
}
func (m *GameState_Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameState_Meta.Merge(m, src)
}
func (m *GameState_Meta) XXX_Size() int {
	return xxx_messageInfo_GameState_Meta.Size(m)
}
func (m *GameState_Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_GameState_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_GameState_Meta proto.InternalMessageInfo

func (m *GameState_Meta) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *GameState_Meta) GetCenterPosition() *Vector {
	if m != nil {
		return m.CenterPosition
	}
	return nil
}

func (m *GameState_Meta) GetScores() []*ScoreEntry {
	if m != nil {
		return m.Scores
	}
	return nil
}

type Entity struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position             *Vector  `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Radius               float64  `protobuf:"fixed64,3,opt,name=radius,proto3" json:"radius,omitempty"`
	Velocity             *Vector  `protobuf:"bytes,4,opt,name=velocity,proto3" json:"velocity,omitempty"`
	Health               int32    `protobuf:"varint,5,opt,name=health,proto3" json:"health,omitempty"`
	BodyDamage           int32    `protobuf:"varint,6,opt,name=body_damage,json=bodyDamage,proto3" json:"body_damage,omitempty"`
	RewardingExperience  int32    `protobuf:"varint,7,opt,name=rewarding_experience,json=rewardingExperience,proto3" json:"rewarding_experience,omitempty"`
	MaxHealth            int32    `protobuf:"varint,8,opt,name=max_health,json=maxHealth,proto3" json:"max_health,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{3}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Entity) GetPosition() *Vector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Entity) GetRadius() float64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *Entity) GetVelocity() *Vector {
	if m != nil {
		return m.Velocity
	}
	return nil
}

func (m *Entity) GetHealth() int32 {
	if m != nil {
		return m.Health
	}
	return 0
}

func (m *Entity) GetBodyDamage() int32 {
	if m != nil {
		return m.BodyDamage
	}
	return 0
}

func (m *Entity) GetRewardingExperience() int32 {
	if m != nil {
		return m.RewardingExperience
	}
	return 0
}

func (m *Entity) GetMaxHealth() int32 {
	if m != nil {
		return m.MaxHealth
	}
	return 0
}

type Vector struct {
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{4}
}

func (m *Vector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector.Unmarshal(m, b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
}
func (m *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(m, src)
}
func (m *Vector) XXX_Size() int {
	return xxx_messageInfo_Vector.Size(m)
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

type ScoreEntry struct {
	HeroId               int32    `protobuf:"varint,1,opt,name=hero_id,json=heroId,proto3" json:"hero_id,omitempty"`
	HeroName             string   `protobuf:"bytes,2,opt,name=hero_name,json=heroName,proto3" json:"hero_name,omitempty"`
	Score                int32    `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	Level                int32    `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScoreEntry) Reset()         { *m = ScoreEntry{} }
func (m *ScoreEntry) String() string { return proto.CompactTextString(m) }
func (*ScoreEntry) ProtoMessage()    {}
func (*ScoreEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{5}
}

func (m *ScoreEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScoreEntry.Unmarshal(m, b)
}
func (m *ScoreEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScoreEntry.Marshal(b, m, deterministic)
}
func (m *ScoreEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScoreEntry.Merge(m, src)
}
func (m *ScoreEntry) XXX_Size() int {
	return xxx_messageInfo_ScoreEntry.Size(m)
}
func (m *ScoreEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ScoreEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ScoreEntry proto.InternalMessageInfo

func (m *ScoreEntry) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *ScoreEntry) GetHeroName() string {
	if m != nil {
		return m.HeroName
	}
	return ""
}

func (m *ScoreEntry) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ScoreEntry) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type Polygon struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Edges                int32    `protobuf:"varint,2,opt,name=edges,proto3" json:"edges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Polygon) Reset()         { *m = Polygon{} }
func (m *Polygon) String() string { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()    {}
func (*Polygon) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{6}
}

func (m *Polygon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Polygon.Unmarshal(m, b)
}
func (m *Polygon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Polygon.Marshal(b, m, deterministic)
}
func (m *Polygon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Polygon.Merge(m, src)
}
func (m *Polygon) XXX_Size() int {
	return xxx_messageInfo_Polygon.Size(m)
}
func (m *Polygon) XXX_DiscardUnknown() {
	xxx_messageInfo_Polygon.DiscardUnknown(m)
}

var xxx_messageInfo_Polygon proto.InternalMessageInfo

func (m *Polygon) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Polygon) GetEdges() int32 {
	if m != nil {
		return m.Edges
	}
	return 0
}

type Bullet struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Owner                int32    `protobuf:"varint,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bullet) Reset()         { *m = Bullet{} }
func (m *Bullet) String() string { return proto.CompactTextString(m) }
func (*Bullet) ProtoMessage()    {}
func (*Bullet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{7}
}

func (m *Bullet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bullet.Unmarshal(m, b)
}
func (m *Bullet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bullet.Marshal(b, m, deterministic)
}
func (m *Bullet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bullet.Merge(m, src)
}
func (m *Bullet) XXX_Size() int {
	return xxx_messageInfo_Bullet.Size(m)
}
func (m *Bullet) XXX_DiscardUnknown() {
	xxx_messageInfo_Bullet.DiscardUnknown(m)
}

var xxx_messageInfo_Bullet proto.InternalMessageInfo

func (m *Bullet) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Bullet) GetOwner() int32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

type Hero struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	AbilityLevels        []int32  `protobuf:"varint,2,rep,packed,name=ability_levels,json=abilityLevels,proto3" json:"ability_levels,omitempty"`
	AbilityValues        []int32  `protobuf:"varint,3,rep,packed,name=ability_values,json=abilityValues,proto3" json:"ability_values,omitempty"`
	SkillPoints          int32    `protobuf:"varint,4,opt,name=skill_points,json=skillPoints,proto3" json:"skill_points,omitempty"`
	Orientation          float64  `protobuf:"fixed64,5,opt,name=orientation,proto3" json:"orientation,omitempty"`
	Level                int32    `protobuf:"varint,6,opt,name=level,proto3" json:"level,omitempty"`
	Score                int32    `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	Experience           int32    `protobuf:"varint,8,opt,name=experience,proto3" json:"experience,omitempty"`
	ExperienceToLevelUp  int32    `protobuf:"varint,9,opt,name=experience_to_level_up,json=experienceToLevelUp,proto3" json:"experience_to_level_up,omitempty"`
	Cooldown             int32    `protobuf:"varint,10,opt,name=cooldown,proto3" json:"cooldown,omitempty"`
	HealthRegenCooldown  int32    `protobuf:"varint,11,opt,name=health_regen_cooldown,json=healthRegenCooldown,proto3" json:"health_regen_cooldown,omitempty"`
	Name                 string   `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hero) Reset()         { *m = Hero{} }
func (m *Hero) String() string { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()    {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{8}
}

func (m *Hero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hero.Unmarshal(m, b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hero.Marshal(b, m, deterministic)
}
func (m *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(m, src)
}
func (m *Hero) XXX_Size() int {
	return xxx_messageInfo_Hero.Size(m)
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *Hero) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Hero) GetAbilityLevels() []int32 {
	if m != nil {
		return m.AbilityLevels
	}
	return nil
}

func (m *Hero) GetAbilityValues() []int32 {
	if m != nil {
		return m.AbilityValues
	}
	return nil
}

func (m *Hero) GetSkillPoints() int32 {
	if m != nil {
		return m.SkillPoints
	}
	return 0
}

func (m *Hero) GetOrientation() float64 {
	if m != nil {
		return m.Orientation
	}
	return 0
}

func (m *Hero) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Hero) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Hero) GetExperience() int32 {
	if m != nil {
		return m.Experience
	}
	return 0
}

func (m *Hero) GetExperienceToLevelUp() int32 {
	if m != nil {
		return m.ExperienceToLevelUp
	}
	return 0
}

func (m *Hero) GetCooldown() int32 {
	if m != nil {
		return m.Cooldown
	}
	return 0
}

func (m *Hero) GetHealthRegenCooldown() int32 {
	if m != nil {
		return m.HealthRegenCooldown
	}
	return 0
}

func (m *Hero) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Command struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Resume               bool     `protobuf:"varint,2,opt,name=resume,proto3" json:"resume,omitempty"`
	Pause                bool     `protobuf:"varint,3,opt,name=pause,proto3" json:"pause,omitempty"`
	Tick                 bool     `protobuf:"varint,4,opt,name=tick,proto3" json:"tick,omitempty"`
	GameReset            bool     `protobuf:"varint,5,opt,name=game_reset,json=gameReset,proto3" json:"game_reset,omitempty"`
	WaitForControls      bool     `protobuf:"varint,6,opt,name=wait_for_controls,json=waitForControls,proto3" json:"wait_for_controls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{9}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Command) GetResume() bool {
	if m != nil {
		return m.Resume
	}
	return false
}

func (m *Command) GetPause() bool {
	if m != nil {
		return m.Pause
	}
	return false
}

func (m *Command) GetTick() bool {
	if m != nil {
		return m.Tick
	}
	return false
}

func (m *Command) GetGameReset() bool {
	if m != nil {
		return m.GameReset
	}
	return false
}

func (m *Command) GetWaitForControls() bool {
	if m != nil {
		return m.WaitForControls
	}
	return false
}

type CommandResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResponse) Reset()         { *m = CommandResponse{} }
func (m *CommandResponse) String() string { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()    {}
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0032e2c8eecaa438, []int{10}
}

func (m *CommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponse.Unmarshal(m, b)
}
func (m *CommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponse.Marshal(b, m, deterministic)
}
func (m *CommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponse.Merge(m, src)
}
func (m *CommandResponse) XXX_Size() int {
	return xxx_messageInfo_CommandResponse.Size(m)
}
func (m *CommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("Ability", Ability_name, Ability_value)
	proto.RegisterType((*Controls)(nil), "Controls")
	proto.RegisterType((*ViewRequest)(nil), "ViewRequest")
	proto.RegisterType((*GameState)(nil), "GameState")
	proto.RegisterType((*GameState_Meta)(nil), "GameState.Meta")
	proto.RegisterType((*Entity)(nil), "Entity")
	proto.RegisterType((*Vector)(nil), "Vector")
	proto.RegisterType((*ScoreEntry)(nil), "ScoreEntry")
	proto.RegisterType((*Polygon)(nil), "Polygon")
	proto.RegisterType((*Bullet)(nil), "Bullet")
	proto.RegisterType((*Hero)(nil), "Hero")
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterType((*CommandResponse)(nil), "CommandResponse")
}

func init() { proto.RegisterFile("thegame.proto", fileDescriptor_0032e2c8eecaa438) }

var fileDescriptor_0032e2c8eecaa438 = []byte{
	// 989 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0xc7, 0x49, 0xfc, 0x27, 0x93, 0x36, 0x49, 0x97, 0xbb, 0x12, 0x15, 0xdd, 0x5d, 0xea, 0x1e,
	0x22, 0xba, 0x07, 0xab, 0xf4, 0xc4, 0x33, 0xdc, 0xf5, 0x0a, 0x87, 0x74, 0xa0, 0x6a, 0xef, 0xe8,
	0x03, 0x2f, 0xd6, 0xd6, 0x1e, 0x12, 0xab, 0xb6, 0xd7, 0xd8, 0x9b, 0xb6, 0xf9, 0x00, 0xbc, 0xf3,
	0x0d, 0xe0, 0x03, 0xf0, 0x85, 0xf8, 0x36, 0x68, 0x67, 0x37, 0x8e, 0x8b, 0x90, 0xe0, 0xc9, 0x9e,
	0xdf, 0x6f, 0x76, 0x76, 0xf7, 0x37, 0xb3, 0x33, 0xb0, 0xaf, 0x56, 0xb8, 0x14, 0x05, 0x46, 0x55,
	0x2d, 0x95, 0x0c, 0xff, 0x72, 0x20, 0x38, 0x97, 0xa5, 0xaa, 0x65, 0xde, 0xb0, 0xa7, 0x00, 0x22,
	0x49, 0x30, 0xc7, 0x5a, 0x28, 0x9c, 0x39, 0x73, 0x67, 0x11, 0xf0, 0x0e, 0xc2, 0xbe, 0x84, 0xc3,
	0xd6, 0xca, 0x64, 0x19, 0xa7, 0x59, 0x8d, 0x89, 0xfe, 0x9b, 0xf5, 0xe6, 0xce, 0xc2, 0xe1, 0x8f,
	0xbb, 0xec, 0x9b, 0x2d, 0xc9, 0x1e, 0x81, 0xdb, 0xac, 0xa4, 0x54, 0xb3, 0x3e, 0x45, 0x34, 0x06,
	0xfb, 0x1c, 0x26, 0xf4, 0xd3, 0x89, 0x32, 0xa0, 0x28, 0x63, 0x82, 0x77, 0xcb, 0x4f, 0x20, 0xc8,
	0xf1, 0x16, 0xf3, 0x78, 0x5d, 0xcd, 0xdc, 0x79, 0x7f, 0x31, 0x3e, 0x0b, 0xa2, 0x57, 0xd7, 0x59,
	0x9e, 0xa9, 0x0d, 0xf7, 0x89, 0xf9, 0xb1, 0x62, 0x0c, 0x06, 0xa5, 0x28, 0x70, 0xe6, 0xcd, 0x9d,
	0xc5, 0x90, 0xd3, 0x7f, 0x78, 0x02, 0xa3, 0xab, 0x0c, 0xef, 0x38, 0xfe, 0xb2, 0xc6, 0x46, 0xe9,
	0x63, 0x28, 0x79, 0x83, 0x25, 0x5d, 0x6c, 0xc8, 0x8d, 0x11, 0xfe, 0xd6, 0x83, 0xe1, 0xb7, 0xa2,
	0xc0, 0xf7, 0x4a, 0xdf, 0xf0, 0x04, 0x06, 0x05, 0x2a, 0x41, 0x2e, 0xa3, 0xb3, 0x49, 0xd4, 0x32,
	0xd1, 0xf7, 0xa8, 0x04, 0x27, 0x92, 0x3d, 0x87, 0xa0, 0x92, 0xf9, 0x66, 0x29, 0xcb, 0x66, 0xd6,
	0x9b, 0xf7, 0x17, 0xa3, 0xb3, 0x20, 0xba, 0x34, 0x00, 0x6f, 0x19, 0x76, 0x0c, 0xfe, 0xf5, 0x3a,
	0xcf, 0x51, 0x35, 0xb3, 0x3e, 0x39, 0xf9, 0xd1, 0x6b, 0xb2, 0xf9, 0x16, 0x67, 0x4f, 0xc0, 0x5b,
	0x61, 0x2d, 0xb1, 0x99, 0x0d, 0xc8, 0xc3, 0x8d, 0xde, 0x62, 0x2d, 0xb9, 0x05, 0x8f, 0x6e, 0x61,
	0xa0, 0x77, 0x65, 0x9f, 0x80, 0xaf, 0x91, 0x38, 0x4b, 0x49, 0x67, 0xd7, 0x38, 0x7c, 0x97, 0xb2,
	0x53, 0x98, 0x24, 0x58, 0x2a, 0xac, 0xe3, 0x4a, 0x36, 0x19, 0x49, 0xe8, 0xd2, 0xc1, 0xfd, 0xe8,
	0x0a, 0x13, 0x25, 0x6b, 0x3e, 0x36, 0xfc, 0xa5, 0xa5, 0xd9, 0x09, 0x78, 0x4d, 0x22, 0xeb, 0x76,
	0xc7, 0x51, 0xf4, 0x5e, 0x9b, 0x17, 0xa5, 0xaa, 0x37, 0xdc, 0x52, 0xe1, 0xaf, 0x3d, 0xf0, 0x2e,
	0x4a, 0x95, 0xa9, 0x0d, 0x1b, 0x43, 0x2f, 0x4b, 0x49, 0x0d, 0x97, 0xf7, 0xb2, 0x54, 0xe7, 0xa2,
	0xdd, 0xaa, 0xf7, 0x70, 0xab, 0x96, 0x60, 0x87, 0xe0, 0xd5, 0x22, 0xcd, 0xd6, 0x0d, 0x25, 0xdc,
	0xe1, 0xd6, 0xd2, 0x8b, 0x6f, 0x31, 0x97, 0x49, 0xa6, 0x36, 0x94, 0xea, 0xee, 0xe2, 0x2d, 0xa1,
	0x17, 0xaf, 0x50, 0xe4, 0x6a, 0x45, 0x57, 0xa1, 0xbb, 0x6a, 0x8b, 0x3d, 0x83, 0xd1, 0xb5, 0x4c,
	0x37, 0x71, 0x2a, 0x0a, 0xb1, 0x34, 0x79, 0x76, 0x39, 0x68, 0xe8, 0x0d, 0x21, 0xec, 0x0b, 0x78,
	0x54, 0xe3, 0x9d, 0xa8, 0xd3, 0xac, 0x5c, 0xc6, 0x78, 0x5f, 0x61, 0x9d, 0x61, 0x99, 0xe0, 0xcc,
	0x27, 0xcf, 0x8f, 0x5b, 0xee, 0xa2, 0xa5, 0xd8, 0x13, 0x80, 0x42, 0xdc, 0xc7, 0x76, 0xbf, 0x80,
	0x1c, 0x87, 0x85, 0xb8, 0x7f, 0x4b, 0x40, 0xf8, 0x1c, 0x3c, 0x73, 0x3c, 0xb6, 0x07, 0xce, 0x3d,
	0xa9, 0xe0, 0x70, 0xe7, 0x5e, 0x5b, 0x1b, 0x5b, 0xf1, 0xce, 0x26, 0x2c, 0x01, 0x76, 0x1a, 0x76,
	0x73, 0xe5, 0x3c, 0xc8, 0xd5, 0xa7, 0x30, 0x24, 0x82, 0xaa, 0xb4, 0x47, 0x15, 0x18, 0x68, 0xe0,
	0x07, 0x51, 0x20, 0xbd, 0x10, 0x1d, 0x83, 0x04, 0x73, 0xb9, 0x31, 0x34, 0x4a, 0xe5, 0x4d, 0x62,
	0xb9, 0xdc, 0x18, 0xe1, 0xd7, 0xe0, 0xdb, 0x62, 0x63, 0xcf, 0xc0, 0x43, 0xca, 0x93, 0xad, 0x57,
	0x3f, 0x32, 0x69, 0xe3, 0x16, 0xd6, 0x11, 0x30, 0x5d, 0x62, 0x63, 0xeb, 0xc6, 0x18, 0xe1, 0x57,
	0xe0, 0x99, 0x4a, 0xfc, 0x5f, 0x01, 0xe4, 0x5d, 0x89, 0xf5, 0x36, 0x00, 0x19, 0xe1, 0xef, 0x7d,
	0x18, 0xe8, 0x4a, 0xfd, 0xef, 0xf5, 0x9f, 0xc1, 0x58, 0x98, 0xa7, 0x1a, 0xd3, 0xe9, 0xcd, 0x83,
	0x71, 0xf9, 0xbe, 0x45, 0xdf, 0x11, 0xd8, 0x75, 0xbb, 0x15, 0xf9, 0x1a, 0xcd, 0x93, 0xd9, 0xb9,
	0x5d, 0x11, 0xc8, 0x8e, 0x61, 0xaf, 0xb9, 0xc9, 0xf2, 0x3c, 0xae, 0x64, 0x56, 0xaa, 0xc6, 0xea,
	0x32, 0x22, 0xec, 0x92, 0x20, 0x36, 0x87, 0x91, 0xd4, 0xd9, 0x55, 0xa2, 0x7d, 0x0e, 0x0e, 0xef,
	0x42, 0x3b, 0x55, 0xbd, 0x8e, 0xaa, 0xbb, 0x0c, 0xf8, 0xdd, 0x0c, 0x3c, 0x05, 0xe8, 0x54, 0x92,
	0x29, 0x90, 0x0e, 0xc2, 0x5e, 0xc2, 0xe1, 0xce, 0x8a, 0x95, 0x8c, 0xdb, 0x46, 0x35, 0x34, 0x55,
	0xb7, 0x63, 0x3f, 0xc8, 0x77, 0xb6, 0x55, 0x1d, 0x41, 0x90, 0x48, 0x99, 0xa7, 0xf2, 0xae, 0x9c,
	0x01, 0xb9, 0xb5, 0x36, 0x3b, 0x83, 0xc7, 0xa6, 0x1a, 0xe3, 0x1a, 0x97, 0x58, 0xc6, 0xad, 0xe3,
	0xc8, 0xc4, 0x33, 0x24, 0xd7, 0xdc, 0xf9, 0x76, 0xcd, 0xb6, 0xf5, 0xed, 0x75, 0x5a, 0xdf, 0x9f,
	0x0e, 0xf8, 0xe7, 0xb2, 0x28, 0x44, 0x99, 0xfe, 0x7b, 0xdf, 0xa3, 0x47, 0x8a, 0xcd, 0xda, 0x16,
	0x63, 0xc0, 0xad, 0xa5, 0xbd, 0x2b, 0xb1, 0x6e, 0x70, 0xdb, 0xac, 0xc9, 0xd0, 0x7b, 0xa8, 0x2c,
	0xb9, 0x21, 0xc5, 0x03, 0x4e, 0xff, 0xfa, 0xf5, 0xe8, 0x41, 0x12, 0xd7, 0xd8, 0xa0, 0x22, 0xa5,
	0x03, 0x3e, 0xd4, 0x08, 0xd7, 0x00, 0x7b, 0x01, 0x07, 0x77, 0x22, 0x53, 0xf1, 0xcf, 0xb2, 0x8e,
	0x13, 0x3b, 0x61, 0x48, 0xf3, 0x80, 0x4f, 0x34, 0xf1, 0x8d, 0xac, 0xb7, 0x83, 0x27, 0x3c, 0x80,
	0x89, 0x3d, 0x2d, 0xc7, 0xa6, 0x92, 0x65, 0x83, 0x2f, 0xfe, 0x70, 0xc0, 0xb7, 0x5d, 0x9e, 0x4d,
	0x61, 0xaf, 0xab, 0xca, 0xf4, 0x23, 0x36, 0xee, 0xbe, 0xdc, 0xa9, 0xc3, 0x26, 0x0f, 0xba, 0xc3,
	0xb4, 0xa7, 0x97, 0x98, 0x2e, 0x1b, 0x37, 0x15, 0x62, 0x3a, 0xed, 0xb3, 0x43, 0x60, 0x16, 0xa9,
	0xb0, 0x44, 0x65, 0x86, 0xd4, 0x74, 0xc0, 0x0e, 0x60, 0xdf, 0xe2, 0x76, 0xb1, 0xdb, 0x81, 0x6a,
	0xcc, 0xa5, 0x48, 0xa7, 0x1e, 0x63, 0x30, 0x2e, 0xe4, 0x2d, 0x16, 0x58, 0x6e, 0x23, 0xfa, 0x67,
	0x6b, 0xf0, 0x3f, 0xac, 0x50, 0x8f, 0x08, 0x76, 0x0c, 0x03, 0xfa, 0x0e, 0xa3, 0xed, 0x9d, 0x8e,
	0x60, 0x37, 0x3c, 0x16, 0xce, 0xa9, 0xc3, 0x42, 0x18, 0xe8, 0x69, 0xc4, 0xf6, 0xa2, 0xce, 0x50,
	0xea, 0x7a, 0x9d, 0x3a, 0xec, 0x04, 0xdc, 0x57, 0x69, 0x91, 0x95, 0x2c, 0x88, 0xac, 0x1e, 0x47,
	0xd3, 0xe8, 0x1f, 0xca, 0xbc, 0x1e, 0xfc, 0xd4, 0xab, 0xae, 0xaf, 0x3d, 0x9a, 0xdf, 0x2f, 0xff,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x78, 0xf6, 0xf5, 0x5e, 0xd0, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TheGameClient is the client API for TheGame service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TheGameClient interface {
	Game(ctx context.Context, opts ...grpc.CallOption) (TheGame_GameClient, error)
	View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (TheGame_ViewClient, error)
	Admin(ctx context.Context, in *Command, opts ...grpc.CallOption) (*CommandResponse, error)
}

type theGameClient struct {
	cc *grpc.ClientConn
}

func NewTheGameClient(cc *grpc.ClientConn) TheGameClient {
	return &theGameClient{cc}
}

func (c *theGameClient) Game(ctx context.Context, opts ...grpc.CallOption) (TheGame_GameClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TheGame_serviceDesc.Streams[0], "/TheGame/Game", opts...)
	if err != nil {
		return nil, err
	}
	x := &theGameGameClient{stream}
	return x, nil
}

type TheGame_GameClient interface {
	Send(*Controls) error
	Recv() (*GameState, error)
	grpc.ClientStream
}

type theGameGameClient struct {
	grpc.ClientStream
}

func (x *theGameGameClient) Send(m *Controls) error {
	return x.ClientStream.SendMsg(m)
}

func (x *theGameGameClient) Recv() (*GameState, error) {
	m := new(GameState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *theGameClient) View(ctx context.Context, in *ViewRequest, opts ...grpc.CallOption) (TheGame_ViewClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TheGame_serviceDesc.Streams[1], "/TheGame/View", opts...)
	if err != nil {
		return nil, err
	}
	x := &theGameViewClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TheGame_ViewClient interface {
	Recv() (*GameState, error)
	grpc.ClientStream
}

type theGameViewClient struct {
	grpc.ClientStream
}

func (x *theGameViewClient) Recv() (*GameState, error) {
	m := new(GameState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *theGameClient) Admin(ctx context.Context, in *Command, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/TheGame/Admin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TheGameServer is the server API for TheGame service.
type TheGameServer interface {
	Game(TheGame_GameServer) error
	View(*ViewRequest, TheGame_ViewServer) error
	Admin(context.Context, *Command) (*CommandResponse, error)
}

func RegisterTheGameServer(s *grpc.Server, srv TheGameServer) {
	s.RegisterService(&_TheGame_serviceDesc, srv)
}

func _TheGame_Game_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TheGameServer).Game(&theGameGameServer{stream})
}

type TheGame_GameServer interface {
	Send(*GameState) error
	Recv() (*Controls, error)
	grpc.ServerStream
}

type theGameGameServer struct {
	grpc.ServerStream
}

func (x *theGameGameServer) Send(m *GameState) error {
	return x.ServerStream.SendMsg(m)
}

func (x *theGameGameServer) Recv() (*Controls, error) {
	m := new(Controls)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TheGame_View_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ViewRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TheGameServer).View(m, &theGameViewServer{stream})
}

type TheGame_ViewServer interface {
	Send(*GameState) error
	grpc.ServerStream
}

type theGameViewServer struct {
	grpc.ServerStream
}

func (x *theGameViewServer) Send(m *GameState) error {
	return x.ServerStream.SendMsg(m)
}

func _TheGame_Admin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TheGameServer).Admin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TheGame/Admin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TheGameServer).Admin(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _TheGame_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TheGame",
	HandlerType: (*TheGameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Admin",
			Handler:    _TheGame_Admin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Game",
			Handler:       _TheGame_Game_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "View",
			Handler:       _TheGame_View_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "thegame.proto",
}
