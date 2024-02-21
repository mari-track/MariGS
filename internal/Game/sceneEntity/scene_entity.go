package sceneEntity

type SceneEntity struct {
	entityIdCounter uint32
	AvatarEntity    map[uint32]*AvatarEntity
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

func (s *SceneEntity) GetNextWorldEntityId(entityType uint8) uint32 {
	for {
		s.entityIdCounter++
		ret := (uint32(entityType) << 24) + s.entityIdCounter
		reTry := false
		for id := range s.AvatarEntity {
			// _, exist := scene.entityMap[ret]
			if ret == id {
				reTry = true
				break
			}
		}
		if reTry {
			continue
		} else {
			return ret
		}
	}
}
