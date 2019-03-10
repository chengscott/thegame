# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: thegame/thegame.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='thegame/thegame.proto',
  package='',
  syntax='proto3',
  serialized_options=_b('Z\002pb'),
  serialized_pb=_b('\n\x15thegame/thegame.proto\"\x90\x01\n\x08\x43ontrols\x12\x12\n\naccelerate\x18\x01 \x01(\x08\x12\x1e\n\x16\x61\x63\x63\x65leration_direction\x18\x02 \x01(\x01\x12\r\n\x05shoot\x18\x03 \x01(\x08\x12\x17\n\x0fshoot_direction\x18\x04 \x01(\x01\x12\x1a\n\x08level_up\x18\x05 \x03(\x0e\x32\x08.Ability\x12\x0c\n\x04name\x18\x06 \x01(\t\"\x1c\n\x0bViewRequest\x12\r\n\x05token\x18\x01 \x01(\t\"\xcf\x01\n\tGameState\x12\x1d\n\x04meta\x18\x01 \x01(\x0b\x32\x0f.GameState.Meta\x12\x1a\n\x08polygons\x18\x02 \x03(\x0b\x32\x08.Polygon\x12\x18\n\x07\x62ullets\x18\x03 \x03(\x0b\x32\x07.Bullet\x12\x15\n\x06heroes\x18\x04 \x03(\x0b\x32\x05.Hero\x1aV\n\x04Meta\x12\x0f\n\x07hero_id\x18\x02 \x01(\x05\x12 \n\x0f\x63\x65nter_position\x18\x05 \x01(\x0b\x32\x07.Vector\x12\x1b\n\x06scores\x18\x04 \x03(\x0b\x32\x0b.ScoreEntry\"\xb1\x01\n\x06\x45ntity\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x19\n\x08position\x18\x02 \x01(\x0b\x32\x07.Vector\x12\x0e\n\x06radius\x18\x03 \x01(\x01\x12\x19\n\x08velocity\x18\x04 \x01(\x0b\x32\x07.Vector\x12\x0e\n\x06health\x18\x05 \x01(\x05\x12\x13\n\x0b\x62ody_damage\x18\x06 \x01(\x05\x12\x1c\n\x14rewarding_experience\x18\x07 \x01(\x05\x12\x12\n\nmax_health\x18\x08 \x01(\x05\"\x1e\n\x06Vector\x12\t\n\x01x\x18\x01 \x01(\x01\x12\t\n\x01y\x18\x02 \x01(\x01\"N\n\nScoreEntry\x12\x0f\n\x07hero_id\x18\x01 \x01(\x05\x12\x11\n\thero_name\x18\x02 \x01(\t\x12\r\n\x05score\x18\x03 \x01(\x05\x12\r\n\x05level\x18\x04 \x01(\x05\"1\n\x07Polygon\x12\x17\n\x06\x65ntity\x18\x01 \x01(\x0b\x32\x07.Entity\x12\r\n\x05\x65\x64ges\x18\x02 \x01(\x05\"0\n\x06\x42ullet\x12\x17\n\x06\x65ntity\x18\x01 \x01(\x0b\x32\x07.Entity\x12\r\n\x05owner\x18\x02 \x01(\x05\"\x8b\x02\n\x04Hero\x12\x17\n\x06\x65ntity\x18\x01 \x01(\x0b\x32\x07.Entity\x12\x16\n\x0e\x61\x62ility_levels\x18\x02 \x03(\x05\x12\x16\n\x0e\x61\x62ility_values\x18\x03 \x03(\x05\x12\x14\n\x0cskill_points\x18\x04 \x01(\x05\x12\x13\n\x0borientation\x18\x05 \x01(\x01\x12\r\n\x05level\x18\x06 \x01(\x05\x12\r\n\x05score\x18\x07 \x01(\x05\x12\x12\n\nexperience\x18\x08 \x01(\x05\x12\x1e\n\x16\x65xperience_to_level_up\x18\t \x01(\x05\x12\x10\n\x08\x63ooldown\x18\n \x01(\x05\x12\x1d\n\x15health_regen_cooldown\x18\x0b \x01(\x05\x12\x0c\n\x04name\x18\x0c \x01(\t\"Y\n\x07\x43ommand\x12\r\n\x05token\x18\x01 \x01(\t\x12\x0e\n\x06resume\x18\x02 \x01(\x08\x12\r\n\x05pause\x18\x03 \x01(\x08\x12\x0c\n\x04tick\x18\x04 \x01(\x08\x12\x12\n\ngame_reset\x18\x05 \x01(\x08\"\x11\n\x0f\x43ommandResponse*\xa0\x01\n\x07\x41\x62ility\x12\x10\n\x0chealth_regen\x10\x00\x12\x0e\n\nmax_health\x10\x01\x12\x0f\n\x0b\x62ody_damage\x10\x02\x12\x10\n\x0c\x62ullet_speed\x10\x03\x12\x16\n\x12\x62ullet_penetration\x10\x04\x12\x11\n\rbullet_damage\x10\x05\x12\x11\n\rbullet_reload\x10\x06\x12\x12\n\x0emovement_speed\x10\x07\x32u\n\x07TheGame\x12!\n\x04Game\x12\t.Controls\x1a\n.GameState(\x01\x30\x01\x12\"\n\x04View\x12\x0c.ViewRequest\x1a\n.GameState0\x01\x12#\n\x05\x41\x64min\x12\x08.Command\x1a\x10.CommandResponseB\x04Z\x02pbb\x06proto3')
)

_ABILITY = _descriptor.EnumDescriptor(
  name='Ability',
  full_name='Ability',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='health_regen', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='max_health', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='body_damage', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='bullet_speed', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='bullet_penetration', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='bullet_damage', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='bullet_reload', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='movement_speed', index=7, number=7,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1186,
  serialized_end=1346,
)
_sym_db.RegisterEnumDescriptor(_ABILITY)

Ability = enum_type_wrapper.EnumTypeWrapper(_ABILITY)
health_regen = 0
max_health = 1
body_damage = 2
bullet_speed = 3
bullet_penetration = 4
bullet_damage = 5
bullet_reload = 6
movement_speed = 7



_CONTROLS = _descriptor.Descriptor(
  name='Controls',
  full_name='Controls',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='accelerate', full_name='Controls.accelerate', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='acceleration_direction', full_name='Controls.acceleration_direction', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='shoot', full_name='Controls.shoot', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='shoot_direction', full_name='Controls.shoot_direction', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='level_up', full_name='Controls.level_up', index=4,
      number=5, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='Controls.name', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=26,
  serialized_end=170,
)


_VIEWREQUEST = _descriptor.Descriptor(
  name='ViewRequest',
  full_name='ViewRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='ViewRequest.token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=172,
  serialized_end=200,
)


_GAMESTATE_META = _descriptor.Descriptor(
  name='Meta',
  full_name='GameState.Meta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hero_id', full_name='GameState.Meta.hero_id', index=0,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='center_position', full_name='GameState.Meta.center_position', index=1,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scores', full_name='GameState.Meta.scores', index=2,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=324,
  serialized_end=410,
)

_GAMESTATE = _descriptor.Descriptor(
  name='GameState',
  full_name='GameState',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='meta', full_name='GameState.meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='polygons', full_name='GameState.polygons', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bullets', full_name='GameState.bullets', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='heroes', full_name='GameState.heroes', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_GAMESTATE_META, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=203,
  serialized_end=410,
)


_ENTITY = _descriptor.Descriptor(
  name='Entity',
  full_name='Entity',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='Entity.id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='position', full_name='Entity.position', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='radius', full_name='Entity.radius', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='velocity', full_name='Entity.velocity', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='health', full_name='Entity.health', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body_damage', full_name='Entity.body_damage', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rewarding_experience', full_name='Entity.rewarding_experience', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_health', full_name='Entity.max_health', index=7,
      number=8, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=413,
  serialized_end=590,
)


_VECTOR = _descriptor.Descriptor(
  name='Vector',
  full_name='Vector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='x', full_name='Vector.x', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='y', full_name='Vector.y', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=592,
  serialized_end=622,
)


_SCOREENTRY = _descriptor.Descriptor(
  name='ScoreEntry',
  full_name='ScoreEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hero_id', full_name='ScoreEntry.hero_id', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hero_name', full_name='ScoreEntry.hero_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='score', full_name='ScoreEntry.score', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='level', full_name='ScoreEntry.level', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=624,
  serialized_end=702,
)


_POLYGON = _descriptor.Descriptor(
  name='Polygon',
  full_name='Polygon',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity', full_name='Polygon.entity', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='edges', full_name='Polygon.edges', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=704,
  serialized_end=753,
)


_BULLET = _descriptor.Descriptor(
  name='Bullet',
  full_name='Bullet',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity', full_name='Bullet.entity', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='owner', full_name='Bullet.owner', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=755,
  serialized_end=803,
)


_HERO = _descriptor.Descriptor(
  name='Hero',
  full_name='Hero',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity', full_name='Hero.entity', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ability_levels', full_name='Hero.ability_levels', index=1,
      number=2, type=5, cpp_type=1, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ability_values', full_name='Hero.ability_values', index=2,
      number=3, type=5, cpp_type=1, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='skill_points', full_name='Hero.skill_points', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='orientation', full_name='Hero.orientation', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='level', full_name='Hero.level', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='score', full_name='Hero.score', index=6,
      number=7, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experience', full_name='Hero.experience', index=7,
      number=8, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experience_to_level_up', full_name='Hero.experience_to_level_up', index=8,
      number=9, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cooldown', full_name='Hero.cooldown', index=9,
      number=10, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='health_regen_cooldown', full_name='Hero.health_regen_cooldown', index=10,
      number=11, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='Hero.name', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=806,
  serialized_end=1073,
)


_COMMAND = _descriptor.Descriptor(
  name='Command',
  full_name='Command',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='token', full_name='Command.token', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resume', full_name='Command.resume', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pause', full_name='Command.pause', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tick', full_name='Command.tick', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='game_reset', full_name='Command.game_reset', index=4,
      number=5, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1075,
  serialized_end=1164,
)


_COMMANDRESPONSE = _descriptor.Descriptor(
  name='CommandResponse',
  full_name='CommandResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1166,
  serialized_end=1183,
)

_CONTROLS.fields_by_name['level_up'].enum_type = _ABILITY
_GAMESTATE_META.fields_by_name['center_position'].message_type = _VECTOR
_GAMESTATE_META.fields_by_name['scores'].message_type = _SCOREENTRY
_GAMESTATE_META.containing_type = _GAMESTATE
_GAMESTATE.fields_by_name['meta'].message_type = _GAMESTATE_META
_GAMESTATE.fields_by_name['polygons'].message_type = _POLYGON
_GAMESTATE.fields_by_name['bullets'].message_type = _BULLET
_GAMESTATE.fields_by_name['heroes'].message_type = _HERO
_ENTITY.fields_by_name['position'].message_type = _VECTOR
_ENTITY.fields_by_name['velocity'].message_type = _VECTOR
_POLYGON.fields_by_name['entity'].message_type = _ENTITY
_BULLET.fields_by_name['entity'].message_type = _ENTITY
_HERO.fields_by_name['entity'].message_type = _ENTITY
DESCRIPTOR.message_types_by_name['Controls'] = _CONTROLS
DESCRIPTOR.message_types_by_name['ViewRequest'] = _VIEWREQUEST
DESCRIPTOR.message_types_by_name['GameState'] = _GAMESTATE
DESCRIPTOR.message_types_by_name['Entity'] = _ENTITY
DESCRIPTOR.message_types_by_name['Vector'] = _VECTOR
DESCRIPTOR.message_types_by_name['ScoreEntry'] = _SCOREENTRY
DESCRIPTOR.message_types_by_name['Polygon'] = _POLYGON
DESCRIPTOR.message_types_by_name['Bullet'] = _BULLET
DESCRIPTOR.message_types_by_name['Hero'] = _HERO
DESCRIPTOR.message_types_by_name['Command'] = _COMMAND
DESCRIPTOR.message_types_by_name['CommandResponse'] = _COMMANDRESPONSE
DESCRIPTOR.enum_types_by_name['Ability'] = _ABILITY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Controls = _reflection.GeneratedProtocolMessageType('Controls', (_message.Message,), dict(
  DESCRIPTOR = _CONTROLS,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Controls)
  ))
_sym_db.RegisterMessage(Controls)

ViewRequest = _reflection.GeneratedProtocolMessageType('ViewRequest', (_message.Message,), dict(
  DESCRIPTOR = _VIEWREQUEST,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:ViewRequest)
  ))
_sym_db.RegisterMessage(ViewRequest)

GameState = _reflection.GeneratedProtocolMessageType('GameState', (_message.Message,), dict(

  Meta = _reflection.GeneratedProtocolMessageType('Meta', (_message.Message,), dict(
    DESCRIPTOR = _GAMESTATE_META,
    __module__ = 'thegame.thegame_pb2'
    # @@protoc_insertion_point(class_scope:GameState.Meta)
    ))
  ,
  DESCRIPTOR = _GAMESTATE,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:GameState)
  ))
_sym_db.RegisterMessage(GameState)
_sym_db.RegisterMessage(GameState.Meta)

Entity = _reflection.GeneratedProtocolMessageType('Entity', (_message.Message,), dict(
  DESCRIPTOR = _ENTITY,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Entity)
  ))
_sym_db.RegisterMessage(Entity)

Vector = _reflection.GeneratedProtocolMessageType('Vector', (_message.Message,), dict(
  DESCRIPTOR = _VECTOR,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Vector)
  ))
_sym_db.RegisterMessage(Vector)

ScoreEntry = _reflection.GeneratedProtocolMessageType('ScoreEntry', (_message.Message,), dict(
  DESCRIPTOR = _SCOREENTRY,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:ScoreEntry)
  ))
_sym_db.RegisterMessage(ScoreEntry)

Polygon = _reflection.GeneratedProtocolMessageType('Polygon', (_message.Message,), dict(
  DESCRIPTOR = _POLYGON,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Polygon)
  ))
_sym_db.RegisterMessage(Polygon)

Bullet = _reflection.GeneratedProtocolMessageType('Bullet', (_message.Message,), dict(
  DESCRIPTOR = _BULLET,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Bullet)
  ))
_sym_db.RegisterMessage(Bullet)

Hero = _reflection.GeneratedProtocolMessageType('Hero', (_message.Message,), dict(
  DESCRIPTOR = _HERO,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Hero)
  ))
_sym_db.RegisterMessage(Hero)

Command = _reflection.GeneratedProtocolMessageType('Command', (_message.Message,), dict(
  DESCRIPTOR = _COMMAND,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:Command)
  ))
_sym_db.RegisterMessage(Command)

CommandResponse = _reflection.GeneratedProtocolMessageType('CommandResponse', (_message.Message,), dict(
  DESCRIPTOR = _COMMANDRESPONSE,
  __module__ = 'thegame.thegame_pb2'
  # @@protoc_insertion_point(class_scope:CommandResponse)
  ))
_sym_db.RegisterMessage(CommandResponse)


DESCRIPTOR._options = None

_THEGAME = _descriptor.ServiceDescriptor(
  name='TheGame',
  full_name='TheGame',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1348,
  serialized_end=1465,
  methods=[
  _descriptor.MethodDescriptor(
    name='Game',
    full_name='TheGame.Game',
    index=0,
    containing_service=None,
    input_type=_CONTROLS,
    output_type=_GAMESTATE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='View',
    full_name='TheGame.View',
    index=1,
    containing_service=None,
    input_type=_VIEWREQUEST,
    output_type=_GAMESTATE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Admin',
    full_name='TheGame.Admin',
    index=2,
    containing_service=None,
    input_type=_COMMAND,
    output_type=_COMMANDRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_THEGAME)

DESCRIPTOR.services_by_name['TheGame'] = _THEGAME

# @@protoc_insertion_point(module_scope)
