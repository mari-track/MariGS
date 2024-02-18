package sceneEntity

type SceneEntity struct {
	AvatarEntity map[uint32]*AvatarEntity
}

type AvatarEntity struct {
	AvatarId       uint32
	WeaponEntityId uint32
}

func (s *SceneEntity) GetAvatarEntity() map[uint32]*AvatarEntity {
	if s.AvatarEntity == nil {
		s.AvatarEntity = make(map[uint32]*AvatarEntity)
	}

	return s.AvatarEntity
}
