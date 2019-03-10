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
	// 963 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4f, 0x6f, 0xdb, 0x36,
	0x14, 0x9f, 0x6c, 0xeb, 0x8f, 0x9f, 0x13, 0xc7, 0xe5, 0xda, 0xcc, 0xc8, 0xd0, 0xd6, 0x51, 0x3a,
	0xcc, 0xd8, 0x41, 0xc8, 0x52, 0xec, 0xbc, 0xb5, 0x69, 0xb0, 0x0e, 0xe8, 0x86, 0x80, 0xed, 0x72,
	0xd8, 0x45, 0x60, 0xac, 0x07, 0x5b, 0x88, 0x44, 0x6a, 0x12, 0x9d, 0xc4, 0xb7, 0x5d, 0x76, 0xdf,
	0x37, 0xd8, 0xbe, 0xd6, 0xbe, 0xcd, 0xc0, 0x47, 0x5a, 0x56, 0x86, 0x01, 0xeb, 0x49, 0xfa, 0xfd,
	0x7e, 0x8f, 0x8f, 0xe4, 0xfb, 0x47, 0xd8, 0xd7, 0x2b, 0x5c, 0x8a, 0x12, 0x93, 0xaa, 0x56, 0x5a,
	0xc5, 0x7f, 0x7b, 0x10, 0x9d, 0x2b, 0xa9, 0x6b, 0x55, 0x34, 0xec, 0x19, 0x80, 0x58, 0x2c, 0xb0,
	0xc0, 0x5a, 0x68, 0x9c, 0x7a, 0x33, 0x6f, 0x1e, 0xf1, 0x0e, 0xc3, 0xbe, 0x81, 0xc3, 0x16, 0xe5,
	0x4a, 0xa6, 0x59, 0x5e, 0xe3, 0xc2, 0xfc, 0x4d, 0x7b, 0x33, 0x6f, 0xee, 0xf1, 0x27, 0x5d, 0xf5,
	0xcd, 0x56, 0x64, 0x8f, 0xc1, 0x6f, 0x56, 0x4a, 0xe9, 0x69, 0x9f, 0x3c, 0x5a, 0xc0, 0xbe, 0x84,
	0x03, 0xfa, 0xe9, 0x78, 0x19, 0x90, 0x97, 0x31, 0xd1, 0xbb, 0xe5, 0x27, 0x10, 0x15, 0x78, 0x8b,
	0x45, 0xba, 0xae, 0xa6, 0xfe, 0xac, 0x3f, 0x1f, 0x9f, 0x45, 0xc9, 0xab, 0xeb, 0xbc, 0xc8, 0xf5,
	0x86, 0x87, 0xa4, 0xfc, 0x5c, 0x31, 0x06, 0x03, 0x29, 0x4a, 0x9c, 0x06, 0x33, 0x6f, 0x3e, 0xe4,
	0xf4, 0x1f, 0x9f, 0xc0, 0xe8, 0x2a, 0xc7, 0x3b, 0x8e, 0xbf, 0xae, 0xb1, 0xd1, 0xe6, 0x18, 0x5a,
	0xdd, 0xa0, 0xa4, 0x8b, 0x0d, 0xb9, 0x05, 0xf1, 0x1f, 0x3d, 0x18, 0x7e, 0x2f, 0x4a, 0x7c, 0xaf,
	0xcd, 0x0d, 0x4f, 0x60, 0x50, 0xa2, 0x16, 0x64, 0x32, 0x3a, 0x3b, 0x48, 0x5a, 0x25, 0xf9, 0x11,
	0xb5, 0xe0, 0x24, 0xb2, 0x17, 0x10, 0x55, 0xaa, 0xd8, 0x2c, 0x95, 0x6c, 0xa6, 0xbd, 0x59, 0x7f,
	0x3e, 0x3a, 0x8b, 0x92, 0x4b, 0x4b, 0xf0, 0x56, 0x61, 0xc7, 0x10, 0x5e, 0xaf, 0x8b, 0x02, 0x75,
	0x33, 0xed, 0x93, 0x51, 0x98, 0xbc, 0x26, 0xcc, 0xb7, 0x3c, 0x7b, 0x0a, 0xc1, 0x0a, 0x6b, 0x85,
	0xcd, 0x74, 0x40, 0x16, 0x7e, 0xf2, 0x16, 0x6b, 0xc5, 0x1d, 0x79, 0x74, 0x0b, 0x03, 0xb3, 0x2b,
	0xfb, 0x0c, 0x42, 0xc3, 0xa4, 0x79, 0x46, 0x71, 0xf6, 0xad, 0xc1, 0x0f, 0x19, 0x3b, 0x85, 0x83,
	0x05, 0x4a, 0x8d, 0x75, 0x5a, 0xa9, 0x26, 0xa7, 0x10, 0xfa, 0x74, 0xf0, 0x30, 0xb9, 0xc2, 0x85,
	0x56, 0x35, 0x1f, 0x5b, 0xfd, 0xd2, 0xc9, 0xec, 0x04, 0x82, 0x66, 0xa1, 0xea, 0x76, 0xc7, 0x51,
	0xf2, 0xde, 0xc0, 0x0b, 0xa9, 0xeb, 0x0d, 0x77, 0x52, 0xfc, 0x7b, 0x0f, 0x82, 0x0b, 0xa9, 0x73,
	0xbd, 0x61, 0x63, 0xe8, 0xe5, 0x19, 0x45, 0xc3, 0xe7, 0xbd, 0x3c, 0x33, 0xb9, 0x68, 0xb7, 0xea,
	0x3d, 0xdc, 0xaa, 0x15, 0xd8, 0x21, 0x04, 0xb5, 0xc8, 0xf2, 0x75, 0x43, 0x09, 0xf7, 0xb8, 0x43,
	0x66, 0xf1, 0x2d, 0x16, 0x6a, 0x91, 0xeb, 0x0d, 0xa5, 0xba, 0xbb, 0x78, 0x2b, 0x98, 0xc5, 0x2b,
	0x14, 0x85, 0x5e, 0xd1, 0x55, 0xe8, 0xae, 0x06, 0xb1, 0xe7, 0x30, 0xba, 0x56, 0xd9, 0x26, 0xcd,
	0x44, 0x29, 0x96, 0x36, 0xcf, 0x3e, 0x07, 0x43, 0xbd, 0x21, 0x86, 0x7d, 0x0d, 0x8f, 0x6b, 0xbc,
	0x13, 0x75, 0x96, 0xcb, 0x65, 0x8a, 0xf7, 0x15, 0xd6, 0x39, 0xca, 0x05, 0x4e, 0x43, 0xb2, 0xfc,
	0xb4, 0xd5, 0x2e, 0x5a, 0x89, 0x3d, 0x05, 0x28, 0xc5, 0x7d, 0xea, 0xf6, 0x8b, 0xc8, 0x70, 0x58,
	0x8a, 0xfb, 0xb7, 0x44, 0xc4, 0x2f, 0x20, 0xb0, 0xc7, 0x63, 0x7b, 0xe0, 0xdd, 0x53, 0x14, 0x3c,
	0xee, 0xdd, 0x1b, 0xb4, 0x71, 0x15, 0xef, 0x6d, 0x62, 0x09, 0xb0, 0x8b, 0x61, 0x37, 0x57, 0xde,
	0x83, 0x5c, 0x7d, 0x0e, 0x43, 0x12, 0xa8, 0x4a, 0x7b, 0x54, 0x81, 0x91, 0x21, 0x7e, 0x12, 0x25,
	0x52, 0x87, 0x18, 0x1f, 0x14, 0x30, 0x9f, 0x5b, 0x60, 0x58, 0x2a, 0x6f, 0x0a, 0x96, 0xcf, 0x2d,
	0x88, 0xbf, 0x83, 0xd0, 0x15, 0x1b, 0x7b, 0x0e, 0x01, 0x52, 0x9e, 0x5c, 0xbd, 0x86, 0x89, 0x4d,
	0x1b, 0x77, 0xb4, 0xf1, 0x80, 0xd9, 0x12, 0x1b, 0x57, 0x37, 0x16, 0xc4, 0xdf, 0x42, 0x60, 0x2b,
	0xf1, 0xa3, 0x1c, 0xa8, 0x3b, 0x89, 0xf5, 0xd6, 0x01, 0x81, 0xf8, 0xcf, 0x3e, 0x0c, 0x4c, 0xa5,
	0xfe, 0xff, 0xfa, 0x2f, 0x60, 0x2c, 0x6c, 0xab, 0xa6, 0x74, 0x7a, 0xdb, 0x30, 0x3e, 0xdf, 0x77,
	0xec, 0x3b, 0x22, 0xbb, 0x66, 0xb7, 0xa2, 0x58, 0xa3, 0x6d, 0x99, 0x9d, 0xd9, 0x15, 0x91, 0xec,
	0x18, 0xf6, 0x9a, 0x9b, 0xbc, 0x28, 0xd2, 0x4a, 0xe5, 0x52, 0x37, 0x2e, 0x2e, 0x23, 0xe2, 0x2e,
	0x89, 0x62, 0x33, 0x18, 0x29, 0x93, 0x5d, 0x2d, 0xda, 0x76, 0xf0, 0x78, 0x97, 0xda, 0x45, 0x35,
	0xe8, 0x44, 0x75, 0x97, 0x81, 0xb0, 0x9b, 0x81, 0x67, 0x00, 0x9d, 0x4a, 0xb2, 0x05, 0xd2, 0x61,
	0xd8, 0x4b, 0x38, 0xdc, 0xa1, 0x54, 0xab, 0xb4, 0x1d, 0x54, 0x43, 0x5b, 0x75, 0x3b, 0xf5, 0x83,
	0x7a, 0xe7, 0x46, 0xd5, 0x11, 0x44, 0x0b, 0xa5, 0x8a, 0x4c, 0xdd, 0xc9, 0x29, 0x90, 0x59, 0x8b,
	0xd9, 0x19, 0x3c, 0xb1, 0xd5, 0x98, 0xd6, 0xb8, 0x44, 0x99, 0xb6, 0x86, 0x23, 0xeb, 0xcf, 0x8a,
	0xdc, 0x68, 0xe7, 0xdb, 0x35, 0xdb, 0xd1, 0xb7, 0xd7, 0x19, 0x7d, 0xbf, 0x79, 0x10, 0x9e, 0xab,
	0xb2, 0x14, 0x32, 0xfb, 0xef, 0xb9, 0x47, 0x4d, 0x8a, 0xcd, 0xda, 0x15, 0x63, 0xc4, 0x1d, 0x32,
	0xd6, 0x95, 0x58, 0x37, 0xb8, 0x1d, 0xd6, 0x04, 0xcc, 0x1e, 0x3a, 0x5f, 0xdc, 0x50, 0xc4, 0x23,
	0x4e, 0xff, 0xa6, 0x7b, 0xcc, 0x43, 0x92, 0xd6, 0xd8, 0xa0, 0xa6, 0x48, 0x47, 0x7c, 0x68, 0x18,
	0x6e, 0x88, 0xf8, 0x11, 0x1c, 0xb8, 0x13, 0x70, 0x6c, 0x2a, 0x25, 0x1b, 0xfc, 0xea, 0x2f, 0x0f,
	0x42, 0x37, 0xb9, 0xd9, 0x04, 0xf6, 0xba, 0x37, 0x9d, 0x7c, 0xc2, 0xc6, 0xdd, 0x6e, 0x9c, 0x78,
	0xec, 0xe0, 0x41, 0xc7, 0x4f, 0x7a, 0x66, 0x89, 0x9d, 0x9c, 0x69, 0x53, 0x21, 0x66, 0x93, 0x3e,
	0x3b, 0x04, 0xe6, 0x98, 0x0a, 0x25, 0x6a, 0xfb, 0xf0, 0x4c, 0x06, 0xec, 0x11, 0xec, 0x3b, 0xde,
	0x2d, 0xf6, 0x3b, 0x54, 0x8d, 0x85, 0x12, 0xd9, 0x24, 0x60, 0x0c, 0xc6, 0xa5, 0xba, 0xc5, 0x12,
	0xe5, 0xd6, 0x63, 0x78, 0xb6, 0x86, 0xf0, 0xc3, 0x0a, 0xcd, 0xd8, 0x67, 0xc7, 0x30, 0xa0, 0xef,
	0x30, 0xd9, 0x3e, 0x90, 0x47, 0xb0, 0x7b, 0x10, 0xe6, 0xde, 0xa9, 0xc7, 0x62, 0x18, 0x98, 0x17,
	0x86, 0xed, 0x25, 0x9d, 0x87, 0xa6, 0x6b, 0x75, 0xea, 0xb1, 0x13, 0xf0, 0x5f, 0x65, 0x65, 0x2e,
	0x59, 0x94, 0xb8, 0x78, 0x1c, 0x4d, 0x92, 0x7f, 0x45, 0xe6, 0xf5, 0xe0, 0x97, 0x5e, 0x75, 0x7d,
	0x1d, 0xd0, 0x9b, 0xfc, 0xf2, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x02, 0x67, 0x85, 0xa4,
	0x07, 0x00, 0x00,
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
