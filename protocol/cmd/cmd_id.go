package cmd

import (
	"github.com/mari-track/MariGS/protocol/proto"
)

const (
	GetShopReq                               = 701
	GetShopRsp                               = 702
	BuyGoodsReq                              = 703
	BuyGoodsRsp                              = 704
	GetShopmallDataReq                       = 705
	GetShopmallDataRsp                       = 706
	OpActivityStateNotify                    = 2501
	SignInInfoReq                            = 2503
	SignInInfoRsp                            = 2504
	GetSignInRewardReq                       = 2505
	GetSignInRewardRsp                       = 2506
	BonusActivityUpdateNotify                = 2512
	BonusActivityInfoReq                     = 2513
	BonusActivityInfoRsp                     = 2514
	GetBonusActivityRewardReq                = 2515
	GetBonusActivityRewardRsp                = 2516
	UnlockAvatarTalentReq                    = 1001
	UnlockAvatarTalentRsp                    = 1002
	AvatarUnlockTalentNotify                 = 1003
	AvatarSkillDepotChangeNotify             = 1004
	BigTalentPointConvertReq                 = 1005
	BigTalentPointConvertRsp                 = 1006
	AvatarSkillMaxChargeCountNotify          = 1007
	AvatarSkillInfoNotify                    = 1008
	ProudSkillUpgradeReq                     = 1009
	ProudSkillUpgradeRsp                     = 1010
	ProudSkillChangeNotify                   = 1011
	AvatarSkillUpgradeReq                    = 1012
	AvatarSkillUpgradeRsp                    = 1013
	AvatarSkillChangeNotify                  = 1014
	ProudSkillExtraLevelNotify               = 1015
	CanUseSkillNotify                        = 1016
	TeamResonanceChangeNotify                = 1017
	GetPlayerFriendListReq                   = 4001
	GetPlayerFriendListRsp                   = 4002
	AskAddFriendReq                          = 4005
	AskAddFriendRsp                          = 4006
	DealAddFriendReq                         = 4007
	DealAddFriendRsp                         = 4008
	GetPlayerSocialDetailReq                 = 4009
	GetPlayerSocialDetailRsp                 = 4010
	DeleteFriendReq                          = 4011
	DeleteFriendRsp                          = 4012
	SetPlayerBirthdayReq                     = 4013
	SetPlayerBirthdayRsp                     = 4014
	SetPlayerSignatureReq                    = 4015
	SetPlayerSignatureRsp                    = 4016
	SetPlayerHeadImageReq                    = 4017
	SetPlayerHeadImageRsp                    = 4018
	UpdatePS4FriendListNotify                = 4019
	DeleteFriendNotify                       = 4020
	AddFriendNotify                          = 4021
	AskAddFriendNotify                       = 4022
	SetNameCardReq                           = 4023
	SetNameCardRsp                           = 4024
	GetAllUnlockNameCardReq                  = 4025
	GetAllUnlockNameCardRsp                  = 4026
	AddBlacklistReq                          = 4027
	AddBlacklistRsp                          = 4028
	RemoveBlacklistReq                       = 4029
	RemoveBlacklistRsp                       = 4030
	UnlockNameCardNotify                     = 4031
	GetRecentMpPlayerListReq                 = 4032
	GetRecentMpPlayerListRsp                 = 4033
	SocialDataNotify                         = 4034
	TakeFirstShareRewardReq                  = 4035
	TakeFirstShareRewardRsp                  = 4036
	UpdatePS4BlockListReq                    = 4037
	UpdatePS4BlockListRsp                    = 4038
	GetPlayerBlacklistReq                    = 4039
	GetPlayerBlacklistRsp                    = 4040
	AbilityInvocationFixedNotify             = 1101
	AbilityInvocationsNotify                 = 1102
	ClientAbilityInitBeginNotify             = 1103
	ClientAbilityInitFinishNotify            = 1104
	AbilityInvocationFailNotify              = 1105
	ClientAbilitiesInitFinishCombineNotify   = 1107
	WindSeedClientNotify                     = 1110
	AbilityChangeNotify                      = 1111
	ClientAbilityChangeNotify                = 1112
	AchievementAllDataNotify                 = 2651
	AchievementUpdateNotify                  = 2652
	TakeAchievementRewardReq                 = 2653
	TakeAchievementRewardRsp                 = 2654
	TakeAchievementGoalRewardReq             = 2655
	TakeAchievementGoalRewardRsp             = 2656
	GetActivityScheduleReq                   = 2001
	GetActivityScheduleRsp                   = 2002
	GetActivityInfoReq                       = 2003
	GetActivityInfoRsp                       = 2004
	ActivityPlayOpenAnimNotify               = 2005
	ActivityInfoNotify                       = 2006
	ActivityScheduleInfoNotify               = 2007
	ActivityTakeWatcherRewardReq             = 2008
	ActivityTakeWatcherRewardRsp             = 2009
	ActivityUpdateWatcherNotify              = 2010
	SeaLampFlyLampReq                        = 2014
	SeaLampFlyLampRsp                        = 2015
	SeaLampTakeContributionRewardReq         = 2016
	SeaLampTakeContributionRewardRsp         = 2017
	SeaLampTakePhaseRewardReq                = 2018
	SeaLampTakePhaseRewardRsp                = 2019
	SeaLampContributeItemReq                 = 2020
	SeaLampContributeItemRsp                 = 2021
	LoadActivityTerrainNotify                = 2024
	ServerAnnounceNotify                     = 2022
	ServerAnnounceRevokeNotify               = 2023
	SalesmanDeliverItemReq                   = 2031
	SalesmanDeliverItemRsp                   = 2032
	SalesmanTakeRewardReq                    = 2033
	SalesmanTakeRewardRsp                    = 2034
	ActivityCondStateChangeNotify            = 2035
	EnterTrialAvatarActivityDungeonReq       = 2041
	EnterTrialAvatarActivityDungeonRsp       = 2042
	ReceivedTrialAvatarActivityRewardReq     = 2043
	ReceivedTrialAvatarActivityRewardRsp     = 2044
	TrialAvatarFirstPassDungeonNotify        = 2047
	TrialAvatarInDungeonIndexNotify          = 2048
	AvatarAddNotify                          = 1701
	AvatarDelNotify                          = 1702
	SetUpAvatarTeamReq                       = 1703
	SetUpAvatarTeamRsp                       = 1704
	ChooseCurAvatarTeamReq                   = 1705
	ChooseCurAvatarTeamRsp                   = 1706
	ChangeAvatarReq                          = 1707
	ChangeAvatarRsp                          = 1708
	AvatarPromoteReq                         = 1709
	AvatarPromoteRsp                         = 1710
	SpringUseReq                             = 1711
	SpringUseRsp                             = 1712
	RefreshBackgroundAvatarReq               = 1713
	RefreshBackgroundAvatarRsp               = 1714
	AvatarTeamUpdateNotify                   = 1715
	AvatarDataNotify                         = 1716
	AvatarUpgradeReq                         = 1717
	AvatarUpgradeRsp                         = 1718
	AvatarDieAnimationEndReq                 = 1719
	AvatarDieAnimationEndRsp                 = 1720
	AvatarChangeElementTypeReq               = 1721
	AvatarChangeElementTypeRsp               = 1722
	AvatarFetterDataNotify                   = 1723
	AvatarExpeditionDataNotify               = 1724
	AvatarExpeditionAllDataReq               = 1725
	AvatarExpeditionAllDataRsp               = 1726
	AvatarExpeditionStartReq                 = 1727
	AvatarExpeditionStartRsp                 = 1728
	AvatarExpeditionCallBackReq              = 1729
	AvatarExpeditionCallBackRsp              = 1730
	AvatarExpeditionGetRewardReq             = 1731
	AvatarExpeditionGetRewardRsp             = 1732
	ChangeMpTeamAvatarReq                    = 1734
	ChangeMpTeamAvatarRsp                    = 1735
	ChangeTeamNameReq                        = 1736
	ChangeTeamNameRsp                        = 1737
	SceneTeamUpdateNotify                    = 1738
	FocusAvatarReq                           = 1740
	FocusAvatarRsp                           = 1741
	AvatarSatiationDataNotify                = 1742
	AvatarWearFlycloakReq                    = 1743
	AvatarWearFlycloakRsp                    = 1744
	AvatarFlycloakChangeNotify               = 1745
	AvatarGainFlycloakNotify                 = 1746
	AvatarEquipAffixStartNotify              = 1747
	AvatarFetterLevelRewardReq               = 1748
	AvatarFetterLevelRewardRsp               = 1749
	BattlePassAllDataNotify                  = 2601
	BattlePassMissionUpdateNotify            = 2602
	BattlePassMissionDelNotify               = 2603
	BattlePassCurScheduleUpdateNotify        = 2604
	TakeBattlePassRewardReq                  = 2605
	TakeBattlePassRewardRsp                  = 2606
	TakeBattlePassMissionPointReq            = 2607
	TakeBattlePassMissionPointRsp            = 2608
	GetBattlePassProductReq                  = 2609
	GetBattlePassProductRsp                  = 2610
	SetBattlePassViewedReq                   = 2613
	SetBattlePassViewedRsp                   = 2614
	BattlePassBuySuccNotify                  = 2615
	BuyBattlePassLevelReq                    = 2616
	BuyBattlePassLevelRsp                    = 2617
	TowerBriefDataNotify                     = 2401
	TowerFloorRecordChangeNotify             = 2402
	TowerCurLevelRecordChangeNotify          = 2403
	TowerDailyRewardProgressChangeNotify     = 2404
	TowerTeamSelectReq                       = 2406
	TowerTeamSelectRsp                       = 2407
	TowerAllDataReq                          = 2408
	TowerAllDataRsp                          = 2409
	TowerEnterLevelReq                       = 2411
	TowerEnterLevelRsp                       = 2412
	TowerBuffSelectReq                       = 2413
	TowerBuffSelectRsp                       = 2414
	TowerSurrenderReq                        = 2421
	TowerSurrenderRsp                        = 2422
	TowerGetFloorStarRewardReq               = 2423
	TowerGetFloorStarRewardRsp               = 2424
	TowerLevelEndNotify                      = 2430
	TowerLevelStarCondNotify                 = 2431
	TowerMiddleLevelChangeTeamNotify         = 2432
	TowerRecordHandbookReq                   = 2433
	TowerRecordHandbookRsp                   = 2434
	WatcherAllDataNotify                     = 2201
	WatcherChangeNotify                      = 2202
	WatcherEventNotify                       = 2203
	WatcherEventTypeNotify                   = 2204
	PushTipsAllDataNotify                    = 2221
	PushTipsChangeNotify                     = 2222
	PushTipsReadFinishReq                    = 2223
	PushTipsReadFinishRsp                    = 2224
	GetPushTipsRewardReq                     = 2225
	GetPushTipsRewardRsp                     = 2226
	GetBlossomBriefInfoListReq               = 2701
	GetBlossomBriefInfoListRsp               = 2702
	BlossomBriefInfoNotify                   = 2703
	WorldOwnerBlossomBriefInfoNotify         = 2704
	WorldOwnerBlossomScheduleInfoNotify      = 2705
	BlossomChestCreateNotify                 = 2706
	OpenBlossomCircleCampGuideNotify         = 2707
	CodexDataFullNotify                      = 4201
	CodexDataUpdataNotify                    = 4202
	QueryCodexMonsterBeKilledNumReq          = 4203
	QueryCodexMonsterBeKilledNumRsp          = 4204
	DungeonEntryInfoReq                      = 901
	DungeonEntryInfoRsp                      = 902
	PlayerEnterDungeonReq                    = 903
	PlayerEnterDungeonRsp                    = 904
	PlayerQuitDungeonReq                     = 905
	PlayerQuitDungeonRsp                     = 906
	DungeonWayPointNotify                    = 907
	DungeonWayPointActivateReq               = 908
	DungeonWayPointActivateRsp               = 909
	DungeonSettleNotify                      = 910
	DungeonPlayerDieNotify                   = 911
	DungeonDieOptionReq                      = 912
	DungeonDieOptionRsp                      = 913
	DungeonShowReminderNotify                = 914
	DungeonPlayerDieReq                      = 915
	DungeonPlayerDieRsp                      = 916
	DungeonDataNotify                        = 917
	DungeonChallengeBeginNotify              = 918
	DungeonChallengeFinishNotify             = 919
	ChallengeDataNotify                      = 920
	DungeonFollowNotify                      = 921
	DungeonGetStatueDropReq                  = 922
	DungeonGetStatueDropRsp                  = 923
	ChallengeRecordNotify                    = 924
	DungeonCandidateTeamInfoNotify           = 925
	DungeonCandidateTeamInviteNotify         = 926
	DungeonCandidateTeamRefuseNotify         = 927
	DungeonCandidateTeamPlayerLeaveNotify    = 928
	DungeonCandidateTeamDismissNotify        = 929
	DungeonCandidateTeamCreateReq            = 930
	DungeonCandidateTeamCreateRsp            = 931
	DungeonCandidateTeamInviteReq            = 932
	DungeonCandidateTeamInviteRsp            = 933
	DungeonCandidateTeamKickReq              = 934
	DungeonCandidateTeamKickRsp              = 935
	DungeonCandidateTeamLeaveReq             = 936
	DungeonCandidateTeamLeaveRsp             = 937
	DungeonCandidateTeamReplyInviteReq       = 938
	DungeonCandidateTeamReplyInviteRsp       = 939
	DungeonCandidateTeamSetReadyReq          = 940
	DungeonCandidateTeamSetReadyRsp          = 941
	DungeonCandidateTeamChangeAvatarReq      = 942
	DungeonCandidateTeamChangeAvatarRsp      = 943
	GetDailyDungeonEntryInfoReq              = 944
	GetDailyDungeonEntryInfoRsp              = 945
	DungeonSlipRevivePointActivateReq        = 946
	DungeonSlipRevivePointActivateRsp        = 947
	DungeonInterruptChallengeReq             = 948
	DungeonInterruptChallengeRsp             = 949
	InteractDailyDungeonInfoNotify           = 950
	EvtBeingHitNotify                        = 301
	EvtAnimatorParameterNotify               = 302
	HostPlayerNotify                         = 303
	EvtDoSkillSuccNotify                     = 304
	EvtCreateGadgetNotify                    = 305
	EvtDestroyGadgetNotify                   = 306
	EvtFaceToEntityNotify                    = 307
	EvtFaceToDirNotify                       = 308
	EvtCostStaminaNotify                     = 309
	EvtSetAttackTargetNotify                 = 310
	EvtAnimatorStateChangedNotify            = 311
	EvtRushMoveNotify                        = 312
	EvtBulletHitNotify                       = 313
	EvtBulletDeactiveNotify                  = 314
	EvtEntityStartDieEndNotify               = 315
	EvtBulletMoveNotify                      = 322
	EvtAvatarEnterFocusNotify                = 323
	EvtAvatarExitFocusNotify                 = 324
	EvtAvatarUpdateFocusNotify               = 325
	EntityAuthorityChangeNotify              = 326
	AvatarBuffAddNotify                      = 327
	AvatarBuffDelNotify                      = 328
	MonsterAlertChangeNotify                 = 329
	MonsterForceAlertNotify                  = 330
	AvatarEnterElementViewNotify             = 332
	TriggerCreateGadgetToEquipPartNotify     = 333
	EvtEntityRenderersChangedNotify          = 334
	AnimatorForceSetAirMoveNotify            = 335
	EvtAiSyncSkillCdNotify                   = 336
	EvtBeingHitsCombineNotify                = 337
	EvtAvatarSitDownNotify                   = 341
	EvtAvatarStandUpNotify                   = 342
	CreateMassiveEntityReq                   = 343
	CreateMassiveEntityRsp                   = 344
	CreateMassiveEntityNotify                = 345
	DestroyMassiveEntityNotify               = 346
	MassiveEntityStateChangedNotify          = 347
	SyncTeamEntityNotify                     = 348
	DelTeamEntityNotify                      = 349
	CombatInvocationsNotify                  = 350
	ServerBuffChangeNotify                   = 351
	EvtAiSyncCombatThreatInfoNotify          = 352
	MassiveEntityElementOpBatchNotify        = 353
	EntityAiSyncNotify                       = 354
	GetGachaInfoReq                          = 1501
	GetGachaInfoRsp                          = 1502
	DoGachaReq                               = 1503
	DoGachaRsp                               = 1504
	GadgetInteractReq                        = 801
	GadgetInteractRsp                        = 802
	GadgetStateNotify                        = 803
	WorktopOptionNotify                      = 804
	SelectWorktopOptionReq                   = 805
	SelectWorktopOptionRsp                   = 806
	BossChestActivateNotify                  = 807
	BlossomChestInfoNotify                   = 808
	GadgetPlayStartNotify                    = 809
	GadgetPlayStopNotify                     = 810
	GadgetPlayDataNotify                     = 811
	GadgetPlayUidOpNotify                    = 812
	PlayerInvestigationAllInfoNotify         = 1901
	TakeInvestigationRewardReq               = 1902
	TakeInvestigationRewardRsp               = 1903
	TakeInvestigationTargetRewardReq         = 1904
	TakeInvestigationTargetRewardRsp         = 1905
	GetInvestigationMonsterReq               = 1906
	GetInvestigationMonsterRsp               = 1907
	PlayerInvestigationNotify                = 1908
	PlayerInvestigationTargetNotify          = 1909
	PlayerStoreNotify                        = 601
	StoreWeightLimitNotify                   = 602
	StoreItemChangeNotify                    = 603
	StoreItemDelNotify                       = 604
	ItemAddHintNotify                        = 605
	UseItemReq                               = 608
	UseItemRsp                               = 609
	DropItemReq                              = 610
	DropItemRsp                              = 611
	WearEquipReq                             = 614
	WearEquipRsp                             = 615
	TakeoffEquipReq                          = 616
	TakeoffEquipRsp                          = 617
	AvatarEquipChangeNotify                  = 618
	WeaponUpgradeReq                         = 619
	WeaponUpgradeRsp                         = 620
	WeaponPromoteReq                         = 621
	WeaponPromoteRsp                         = 622
	ReliquaryUpgradeReq                      = 623
	ReliquaryUpgradeRsp                      = 624
	ReliquaryPromoteReq                      = 625
	ReliquaryPromoteRsp                      = 626
	AvatarCardChangeReq                      = 627
	AvatarCardChangeRsp                      = 628
	GrantRewardNotify                        = 629
	WeaponAwakenReq                          = 630
	WeaponAwakenRsp                          = 631
	ItemCdGroupTimeNotify                    = 632
	DropHintNotify                           = 633
	CombineReq                               = 634
	CombineRsp                               = 635
	ForgeQueueDataNotify                     = 636
	ForgeGetQueueDataReq                     = 637
	ForgeGetQueueDataRsp                     = 638
	ForgeStartReq                            = 639
	ForgeStartRsp                            = 640
	ForgeQueueManipulateReq                  = 641
	ForgeQueueManipulateRsp                  = 642
	ResinChangeNotify                        = 643
	BuyResinReq                              = 649
	BuyResinRsp                              = 650
	MaterialDeleteReturnNotify               = 651
	TakeMaterialDeleteReturnReq              = 652
	TakeMaterialDeleteReturnRsp              = 653
	MaterialDeleteUpdateNotify               = 654
	McoinExchangeHcoinReq                    = 655
	McoinExchangeHcoinRsp                    = 656
	DestroyMaterialReq                       = 657
	DestroyMaterialRsp                       = 658
	MailChangeNotify                         = 1402
	ReadMailNotify                           = 1403
	GetMailItemReq                           = 1404
	GetMailItemRsp                           = 1405
	DelMailReq                               = 1406
	DelMailRsp                               = 1407
	GetAuthkeyReq                            = 1408
	GetAuthkeyRsp                            = 1409
	ClientNewMailNotify                      = 1410
	GetAllMailReq                            = 1411
	GetAllMailRsp                            = 1412
	PlayerStartMatchReq                      = 4151
	PlayerStartMatchRsp                      = 4152
	PlayerMatchInfoNotify                    = 4153
	PlayerCancelMatchReq                     = 4154
	PlayerCancelMatchRsp                     = 4155
	PlayerMatchStopNotify                    = 4156
	PlayerMatchSuccNotify                    = 4157
	PlayerConfirmMatchReq                    = 4158
	PlayerConfirmMatchRsp                    = 4159
	PlayerAllowEnterMpAfterAgreeMatchNotify  = 4160
	PlayerMatchAgreedResultNotify            = 4161
	PlayerApplyEnterMpAfterMatchAgreedNotify = 4162
	KeepAliveNotify                          = 1
	GmTalkReq                                = 2
	GmTalkRsp                                = 3
	ShowMessageNotify                        = 4
	PingReq                                  = 5
	PingRsp                                  = 6
	GetOnlinePlayerListReq                   = 8
	GetOnlinePlayerListRsp                   = 9
	ServerTimeNotify                         = 10
	ServerLogNotify                          = 11
	ClientReconnectNotify                    = 12
	RobotPushPlayerDataNotify                = 14
	ClientReportNotify                       = 15
	UnionCmdNotify                           = 16
	GetOnlinePlayerInfoReq                   = 17
	GetOnlinePlayerInfoRsp                   = 18
	CheckSegmentCRCNotify                    = 19
	CheckSegmentCRCReq                       = 20
	WorldPlayerRTTNotify                     = 21
	EchoNotify                               = 22
	MonsterSummonTagNotify                   = 1301
	PlayerApplyEnterMpNotify                 = 1801
	PlayerApplyEnterMpReq                    = 1802
	PlayerApplyEnterMpRsp                    = 1803
	PlayerApplyEnterMpResultNotify           = 1804
	PlayerApplyEnterMpResultReq              = 1805
	PlayerApplyEnterMpResultRsp              = 1806
	PlayerQuitFromMpNotify                   = 1807
	PlayerPreEnterMpNotify                   = 1808
	GetPlayerMpModeAvailabilityReq           = 1809
	GetPlayerMpModeAvailabilityRsp           = 1810
	PlayerSetOnlyMPWithPSPlayerReq           = 1811
	PlayerSetOnlyMPWithPSPlayerRsp           = 1812
	PSPlayerApplyEnterMpReq                  = 1813
	PSPlayerApplyEnterMpRsp                  = 1814
	MpPlayOwnerCheckReq                      = 1815
	MpPlayOwnerCheckRsp                      = 1816
	MpPlayOwnerStartInviteReq                = 1817
	MpPlayOwnerStartInviteRsp                = 1818
	MpPlayOwnerInviteNotify                  = 1819
	MpPlayGuestReplyInviteReq                = 1820
	MpPlayGuestReplyInviteRsp                = 1821
	MpPlayGuestReplyNotify                   = 1822
	MpPlayPrepareNotify                      = 1823
	MpPlayInviteResultNotify                 = 1824
	MpPlayPrepareInterruptNotify             = 1825
	NpcTalkReq                               = 501
	NpcTalkRsp                               = 502
	GetSceneNpcPositionReq                   = 504
	GetSceneNpcPositionRsp                   = 505
	QueryPathReq                             = 2301
	QueryPathRsp                             = 2302
	ObstacleModifyNotify                     = 2303
	PathfindingPingNotify                    = 2304
	PathfindingEnterSceneReq                 = 2305
	PathfindingEnterSceneRsp                 = 2306
	GMShowObstacleReq                        = 2351
	GMShowObstacleRsp                        = 2352
	GMShowNavMeshReq                         = 2353
	GMShowNavMeshRsp                         = 2354
	NavMeshStatsNotify                       = 2355
	GetPlayerTokenReq                        = 101
	GetPlayerTokenRsp                        = 102
	PlayerLoginReq                           = 103
	PlayerLoginRsp                           = 104
	PlayerLogoutReq                          = 105
	PlayerLogoutRsp                          = 106
	PlayerLogoutNotify                       = 107
	PlayerDataNotify                         = 108
	ChangeGameTimeReq                        = 109
	ChangeGameTimeRsp                        = 110
	PlayerGameTimeNotify                     = 111
	PlayerPropNotify                         = 112
	ClientTriggerEventNotify                 = 113
	SetPlayerPropReq                         = 114
	SetPlayerPropRsp                         = 115
	SetPlayerBornDataReq                     = 116
	SetPlayerBornDataRsp                     = 117
	DoSetPlayerBornDataNotify                = 118
	PlayerPropChangeNotify                   = 119
	SetPlayerNameReq                         = 120
	SetPlayerNameRsp                         = 121
	SetOpenStateReq                          = 122
	SetOpenStateRsp                          = 123
	OpenStateUpdateNotify                    = 124
	OpenStateChangeNotify                    = 125
	PlayerCookReq                            = 126
	PlayerCookRsp                            = 127
	PlayerRandomCookReq                      = 128
	PlayerRandomCookRsp                      = 129
	CookDataNotify                           = 130
	CookRecipeDataNotify                     = 131
	CookGradeDataNotify                      = 132
	PlayerCompoundMaterialReq                = 133
	PlayerCompoundMaterialRsp                = 134
	TakeCompoundOutputReq                    = 135
	TakeCompoundOutputRsp                    = 136
	CompoundDataNotify                       = 137
	GetCompoundDataReq                       = 138
	GetCompoundDataRsp                       = 139
	PlayerTimeNotify                         = 140
	PlayerSetPauseReq                        = 141
	PlayerSetPauseRsp                        = 142
	PlayerSetLanguageReq                     = 143
	PlayerSetLanguageRsp                     = 144
	DataResVersionNotify                     = 145
	DailyTaskDataNotify                      = 146
	DailyTaskProgressNotify                  = 147
	DailyTaskScoreRewardNotify               = 148
	WorldOwnerDailyTaskNotify                = 149
	AddRandTaskInfoNotify                    = 150
	RemoveRandTaskInfoNotify                 = 151
	TakePlayerLevelRewardReq                 = 152
	TakePlayerLevelRewardRsp                 = 153
	PlayerLevelRewardUpdateNotify            = 154
	GivingRecordNotify                       = 155
	GivingRecordChangeNotify                 = 156
	ItemGivingReq                            = 157
	ItemGivingRsp                            = 158
	PlayerCookArgsReq                        = 159
	PlayerCookArgsRsp                        = 160
	PlayerLuaShellNotify                     = 161
	ServerDisconnectClientNotify             = 162
	AntiAddictNotify                         = 163
	PlayerForceExitReq                       = 164
	PlayerForceExitRsp                       = 165
	PlayerInjectFixNotify                    = 166
	TaskVarNotify                            = 167
	ClientLockGameTimeNotify                 = 168
	EntityPropNotify                         = 1201
	LifeStateChangeNotify                    = 1202
	EntityFightPropNotify                    = 1203
	EntityFightPropUpdateNotify              = 1204
	AvatarFightPropNotify                    = 1205
	AvatarFightPropUpdateNotify              = 1206
	EntityFightPropChangeReasonNotify        = 1207
	AvatarLifeStateChangeNotify              = 1208
	AvatarPropChangeReasonNotify             = 1209
	PlayerPropChangeReasonNotify             = 1210
	AvatarPropNotify                         = 1211
	MarkNewNotify                            = 1212
	QuestListNotify                          = 401
	QuestListUpdateNotify                    = 402
	QuestDelNotify                           = 403
	FinishedParentQuestNotify                = 404
	FinishedParentQuestUpdateNotify          = 405
	AddQuestContentProgressReq               = 406
	AddQuestContentProgressRsp               = 407
	GetQuestTalkHistoryReq                   = 408
	GetQuestTalkHistoryRsp                   = 409
	QuestCreateEntityReq                     = 410
	QuestCreateEntityRsp                     = 411
	QuestDestroyEntityReq                    = 412
	QuestDestroyEntityRsp                    = 413
	ChapterStateNotify                       = 416
	QuestProgressUpdateNotify                = 417
	QuestUpdateQuestVarReq                   = 418
	QuestUpdateQuestVarRsp                   = 419
	QuestUpdateQuestVarNotify                = 420
	QuestDestroyNpcReq                       = 421
	QuestDestroyNpcRsp                       = 422
	BargainStartNotify                       = 423
	BargainOfferPriceReq                     = 424
	BargainOfferPriceRsp                     = 425
	BargainTerminateNotify                   = 426
	GetBargainDataReq                        = 427
	GetBargainDataRsp                        = 428
	GetAllActivatedBargainDataReq            = 429
	GetAllActivatedBargainDataRsp            = 430
	ServerCondMeetQuestListUpdateNotify      = 431
	QuestGlobalVarNotify                     = 432
	QuestTransmitReq                         = 433
	QuestTransmitRsp                         = 434
	PersonalLineAllDataReq                   = 435
	PersonalLineAllDataRsp                   = 436
	RedeemLegendaryKeyReq                    = 437
	RedeemLegendaryKeyRsp                    = 438
	UnlockPersonalLineReq                    = 439
	UnlockPersonalLineRsp                    = 440
	RechargeReq                              = 4101
	RechargeRsp                              = 4102
	OrderFinishNotify                        = 4103
	CardProductRewardNotify                  = 4104
	PlayerRechargeDataNotify                 = 4105
	OrderDisplayNotify                       = 4106
	PlayerEnterSceneNotify                   = 201
	LeaveSceneReq                            = 202
	LeaveSceneRsp                            = 203
	SceneInitFinishReq                       = 204
	SceneInitFinishRsp                       = 205
	SceneEntityAppearNotify                  = 206
	SceneEntityDisappearNotify               = 207
	SceneEntityMoveReq                       = 208
	SceneEntityMoveRsp                       = 209
	SceneAvatarStaminaStepReq                = 210
	SceneAvatarStaminaStepRsp                = 211
	SceneEntityMoveNotify                    = 212
	ScenePlayerLocationNotify                = 213
	GetScenePointReq                         = 214
	GetScenePointRsp                         = 215
	EnterTransPointRegionNotify              = 216
	ExitTransPointRegionNotify               = 217
	ScenePointUnlockNotify                   = 218
	SceneTransToPointReq                     = 219
	SceneTransToPointRsp                     = 220
	EntityJumpNotify                         = 221
	GetSceneAreaReq                          = 222
	GetSceneAreaRsp                          = 223
	SceneAreaUnlockNotify                    = 224
	SceneEntityDrownReq                      = 225
	SceneEntityDrownRsp                      = 226
	SceneCreateEntityReq                     = 227
	SceneCreateEntityRsp                     = 228
	SceneDestroyEntityReq                    = 229
	SceneDestroyEntityRsp                    = 230
	SceneForceUnlockNotify                   = 231
	SceneForceLockNotify                     = 232
	EnterWorldAreaReq                        = 233
	EnterWorldAreaRsp                        = 234
	EntityForceSyncReq                       = 235
	EntityForceSyncRsp                       = 236
	SceneAreaExploreNotify                   = 237
	SceneGetAreaExplorePercentReq            = 238
	SceneGetAreaExplorePercentRsp            = 239
	ClientTransmitReq                        = 240
	ClientTransmitRsp                        = 241
	EnterSceneWeatherAreaNotify              = 242
	ExitSceneWeatherAreaNotify               = 243
	SceneAreaWeatherNotify                   = 244
	ScenePlayerInfoNotify                    = 245
	WorldPlayerLocationNotify                = 246
	BeginCameraSceneLookNotify               = 247
	EndCameraSceneLookNotify                 = 248
	MarkEntityInMinMapNotify                 = 249
	UnmarkEntityInMinMapNotify               = 250
	DropSubfieldReq                          = 251
	DropSubfieldRsp                          = 252
	ExecuteGroupTriggerReq                   = 253
	ExecuteGroupTriggerRsp                   = 254
	LevelupCityReq                           = 255
	LevelupCityRsp                           = 256
	SceneRouteChangeNotify                   = 257
	PlatformStartRouteNotify                 = 258
	PlatformStopRouteNotify                  = 259
	PlatformChangeRouteNotify                = 260
	ScenePlayerSoundNotify                   = 261
	PersonalSceneJumpReq                     = 262
	PersonalSceneJumpRsp                     = 263
	SealBattleBeginNotify                    = 264
	SealBattleEndNotify                      = 265
	SealBattleProgressNotify                 = 266
	ClientPauseNotify                        = 267
	PlayerEnterSceneInfoNotify               = 268
	JoinPlayerSceneReq                       = 269
	JoinPlayerSceneRsp                       = 270
	SceneKickPlayerReq                       = 271
	SceneKickPlayerRsp                       = 272
	SceneKickPlayerNotify                    = 273
	HitClientTrivialNotify                   = 274
	BackMyWorldReq                           = 275
	BackMyWorldRsp                           = 276
	SeeMonsterReq                            = 277
	SeeMonsterRsp                            = 278
	AddSeenMonsterNotify                     = 279
	AllSeenMonsterNotify                     = 280
	SceneTimeNotify                          = 281
	EnterSceneReadyReq                       = 282
	EnterSceneReadyRsp                       = 283
	EnterScenePeerNotify                     = 284
	EnterSceneDoneReq                        = 285
	EnterSceneDoneRsp                        = 286
	WorldPlayerDieNotify                     = 287
	WorldPlayerReviveReq                     = 288
	WorldPlayerReviveRsp                     = 289
	JoinPlayerFailNotify                     = 290
	SetSceneWeatherAreaReq                   = 291
	SetSceneWeatherAreaRsp                   = 292
	ExecuteGadgetLuaReq                      = 293
	ExecuteGadgetLuaRsp                      = 294
	CutSceneBeginNotify                      = 295
	CutSceneFinishNotify                     = 296
	CutSceneEndNotify                        = 297
	ClientScriptEventNotify                  = 298
	SceneEntitiesMovesReq                    = 299
	SceneEntitiesMovesRsp                    = 300
	SceneEntitiesMoveCombineNotify           = 3001
	UnlockTransPointReq                      = 3002
	UnlockTransPointRsp                      = 3003
	PlatformRouteStateNotify                 = 3004
	SceneWeatherForcastReq                   = 3005
	SceneWeatherForcastRsp                   = 3006
	MarkMapReq                               = 3010
	MarkMapRsp                               = 3011
	AllMarkPointNotify                       = 3012
	WorldDataNotify                          = 3013
	EntityMoveRoomNotify                     = 3014
	WorldPlayerInfoNotify                    = 3015
	PostEnterSceneReq                        = 3016
	PostEnterSceneRsp                        = 3017
	PlayerChatReq                            = 3018
	PlayerChatRsp                            = 3019
	PlayerChatNotify                         = 3020
	PlayerChatCDNotify                       = 3021
	ChatHistoryNotify                        = 3022
	SceneDataNotify                          = 3023
	DungeonEntryToBeExploreNotify            = 3024
	GetDungeonEntryExploreConditionReq       = 3035
	GetDungeonEntryExploreConditionRsp       = 3036
	UnfreezeGroupLimitNotify                 = 3037
	SetEntityClientDataNotify                = 3041
	GroupSuiteNotify                         = 3042
	GroupUnloadNotify                        = 3043
	MonsterAIConfigHashNotify                = 3044
	EntityAbilityConfigHashNotify            = 3045
	ShowTemplateReminderNotify               = 3046
	ShowCommonTipsNotify                     = 3047
	CloseCommonTipsNotify                    = 3048
)

func (c *CmdProtoMap) registerAllMessage() {
	c.regMsg(GetPlayerTokenReq, func() any { return new(proto.GetPlayerTokenReq) })
	c.regMsg(GetShopReq, func() any { return new(proto.GetShopReq) })
	c.regMsg(GetShopRsp, func() any { return new(proto.GetShopRsp) })
	c.regMsg(BuyGoodsReq, func() any { return new(proto.BuyGoodsReq) })
	c.regMsg(BuyGoodsRsp, func() any { return new(proto.BuyGoodsRsp) })
	c.regMsg(GetShopmallDataReq, func() any { return new(proto.GetShopmallDataReq) })
	c.regMsg(GetShopmallDataRsp, func() any { return new(proto.GetShopmallDataRsp) })
	c.regMsg(OpActivityStateNotify, func() any { return new(proto.OpActivityStateNotify) })
	c.regMsg(SignInInfoReq, func() any { return new(proto.SignInInfoReq) })
	c.regMsg(SignInInfoRsp, func() any { return new(proto.SignInInfoRsp) })
	c.regMsg(GetSignInRewardReq, func() any { return new(proto.GetSignInRewardReq) })
	c.regMsg(GetSignInRewardRsp, func() any { return new(proto.GetSignInRewardRsp) })
	c.regMsg(BonusActivityUpdateNotify, func() any { return new(proto.BonusActivityUpdateNotify) })
	c.regMsg(BonusActivityInfoReq, func() any { return new(proto.BonusActivityInfoReq) })
	c.regMsg(BonusActivityInfoRsp, func() any { return new(proto.BonusActivityInfoRsp) })
	c.regMsg(GetBonusActivityRewardReq, func() any { return new(proto.GetBonusActivityRewardReq) })
	c.regMsg(GetBonusActivityRewardRsp, func() any { return new(proto.GetBonusActivityRewardRsp) })
	c.regMsg(UnlockAvatarTalentReq, func() any { return new(proto.UnlockAvatarTalentReq) })
	c.regMsg(UnlockAvatarTalentRsp, func() any { return new(proto.UnlockAvatarTalentRsp) })
	c.regMsg(AvatarUnlockTalentNotify, func() any { return new(proto.AvatarUnlockTalentNotify) })
	c.regMsg(AvatarSkillDepotChangeNotify, func() any { return new(proto.AvatarSkillDepotChangeNotify) })
	c.regMsg(BigTalentPointConvertReq, func() any { return new(proto.BigTalentPointConvertReq) })
	c.regMsg(BigTalentPointConvertRsp, func() any { return new(proto.BigTalentPointConvertRsp) })
	c.regMsg(AvatarSkillMaxChargeCountNotify, func() any { return new(proto.AvatarSkillMaxChargeCountNotify) })
	c.regMsg(AvatarSkillInfoNotify, func() any { return new(proto.AvatarSkillInfoNotify) })
	c.regMsg(ProudSkillUpgradeReq, func() any { return new(proto.ProudSkillUpgradeReq) })
	c.regMsg(ProudSkillUpgradeRsp, func() any { return new(proto.ProudSkillUpgradeRsp) })
	c.regMsg(ProudSkillChangeNotify, func() any { return new(proto.ProudSkillChangeNotify) })
	c.regMsg(AvatarSkillUpgradeReq, func() any { return new(proto.AvatarSkillUpgradeReq) })
	c.regMsg(AvatarSkillUpgradeRsp, func() any { return new(proto.AvatarSkillUpgradeRsp) })
	c.regMsg(AvatarSkillChangeNotify, func() any { return new(proto.AvatarSkillChangeNotify) })
	c.regMsg(ProudSkillExtraLevelNotify, func() any { return new(proto.ProudSkillExtraLevelNotify) })
	c.regMsg(CanUseSkillNotify, func() any { return new(proto.CanUseSkillNotify) })
	c.regMsg(TeamResonanceChangeNotify, func() any { return new(proto.TeamResonanceChangeNotify) })
	c.regMsg(GetPlayerFriendListReq, func() any { return new(proto.GetPlayerFriendListReq) })
	c.regMsg(GetPlayerFriendListRsp, func() any { return new(proto.GetPlayerFriendListRsp) })
	c.regMsg(AskAddFriendReq, func() any { return new(proto.AskAddFriendReq) })
	c.regMsg(AskAddFriendRsp, func() any { return new(proto.AskAddFriendRsp) })
	c.regMsg(DealAddFriendReq, func() any { return new(proto.DealAddFriendReq) })
	c.regMsg(DealAddFriendRsp, func() any { return new(proto.DealAddFriendRsp) })
	c.regMsg(GetPlayerSocialDetailReq, func() any { return new(proto.GetPlayerSocialDetailReq) })
	c.regMsg(GetPlayerSocialDetailRsp, func() any { return new(proto.GetPlayerSocialDetailRsp) })
	c.regMsg(DeleteFriendReq, func() any { return new(proto.DeleteFriendReq) })
	c.regMsg(DeleteFriendRsp, func() any { return new(proto.DeleteFriendRsp) })
	c.regMsg(SetPlayerBirthdayReq, func() any { return new(proto.SetPlayerBirthdayReq) })
	c.regMsg(SetPlayerBirthdayRsp, func() any { return new(proto.SetPlayerBirthdayRsp) })
	c.regMsg(SetPlayerSignatureReq, func() any { return new(proto.SetPlayerSignatureReq) })
	c.regMsg(SetPlayerSignatureRsp, func() any { return new(proto.SetPlayerSignatureRsp) })
	c.regMsg(SetPlayerHeadImageReq, func() any { return new(proto.SetPlayerHeadImageReq) })
	c.regMsg(SetPlayerHeadImageRsp, func() any { return new(proto.SetPlayerHeadImageRsp) })
	c.regMsg(UpdatePS4FriendListNotify, func() any { return new(proto.UpdatePS4FriendListNotify) })
	c.regMsg(DeleteFriendNotify, func() any { return new(proto.DeleteFriendNotify) })
	c.regMsg(AddFriendNotify, func() any { return new(proto.AddFriendNotify) })
	c.regMsg(AskAddFriendNotify, func() any { return new(proto.AskAddFriendNotify) })
	c.regMsg(SetNameCardReq, func() any { return new(proto.SetNameCardReq) })
	c.regMsg(SetNameCardRsp, func() any { return new(proto.SetNameCardRsp) })
	c.regMsg(GetAllUnlockNameCardReq, func() any { return new(proto.GetAllUnlockNameCardReq) })
	c.regMsg(GetAllUnlockNameCardRsp, func() any { return new(proto.GetAllUnlockNameCardRsp) })
	c.regMsg(AddBlacklistReq, func() any { return new(proto.AddBlacklistReq) })
	c.regMsg(AddBlacklistRsp, func() any { return new(proto.AddBlacklistRsp) })
	c.regMsg(RemoveBlacklistReq, func() any { return new(proto.RemoveBlacklistReq) })
	c.regMsg(RemoveBlacklistRsp, func() any { return new(proto.RemoveBlacklistRsp) })
	c.regMsg(UnlockNameCardNotify, func() any { return new(proto.UnlockNameCardNotify) })
	c.regMsg(GetRecentMpPlayerListReq, func() any { return new(proto.GetRecentMpPlayerListReq) })
	c.regMsg(GetRecentMpPlayerListRsp, func() any { return new(proto.GetRecentMpPlayerListRsp) })
	c.regMsg(SocialDataNotify, func() any { return new(proto.SocialDataNotify) })
	c.regMsg(TakeFirstShareRewardReq, func() any { return new(proto.TakeFirstShareRewardReq) })
	c.regMsg(TakeFirstShareRewardRsp, func() any { return new(proto.TakeFirstShareRewardRsp) })
	c.regMsg(UpdatePS4BlockListReq, func() any { return new(proto.UpdatePS4BlockListReq) })
	c.regMsg(UpdatePS4BlockListRsp, func() any { return new(proto.UpdatePS4BlockListRsp) })
	c.regMsg(GetPlayerBlacklistReq, func() any { return new(proto.GetPlayerBlacklistReq) })
	c.regMsg(GetPlayerBlacklistRsp, func() any { return new(proto.GetPlayerBlacklistRsp) })
	c.regMsg(AbilityInvocationFixedNotify, func() any { return new(proto.AbilityInvocationFixedNotify) })
	c.regMsg(AbilityInvocationsNotify, func() any { return new(proto.AbilityInvocationsNotify) })
	c.regMsg(ClientAbilityInitBeginNotify, func() any { return new(proto.ClientAbilityInitBeginNotify) })
	c.regMsg(ClientAbilityInitFinishNotify, func() any { return new(proto.ClientAbilityInitFinishNotify) })
	c.regMsg(AbilityInvocationFailNotify, func() any { return new(proto.AbilityInvocationFailNotify) })
	c.regMsg(ClientAbilitiesInitFinishCombineNotify, func() any { return new(proto.ClientAbilitiesInitFinishCombineNotify) })
	c.regMsg(WindSeedClientNotify, func() any { return new(proto.WindSeedClientNotify) })
	c.regMsg(AbilityChangeNotify, func() any { return new(proto.AbilityChangeNotify) })
	c.regMsg(ClientAbilityChangeNotify, func() any { return new(proto.ClientAbilityChangeNotify) })
	c.regMsg(AchievementAllDataNotify, func() any { return new(proto.AchievementAllDataNotify) })
	c.regMsg(AchievementUpdateNotify, func() any { return new(proto.AchievementUpdateNotify) })
	c.regMsg(TakeAchievementRewardReq, func() any { return new(proto.TakeAchievementRewardReq) })
	c.regMsg(TakeAchievementRewardRsp, func() any { return new(proto.TakeAchievementRewardRsp) })
	c.regMsg(TakeAchievementGoalRewardReq, func() any { return new(proto.TakeAchievementGoalRewardReq) })
	c.regMsg(TakeAchievementGoalRewardRsp, func() any { return new(proto.TakeAchievementGoalRewardRsp) })
	c.regMsg(GetActivityScheduleReq, func() any { return new(proto.GetActivityScheduleReq) })
	c.regMsg(GetActivityScheduleRsp, func() any { return new(proto.GetActivityScheduleRsp) })
	c.regMsg(GetActivityInfoReq, func() any { return new(proto.GetActivityInfoReq) })
	c.regMsg(GetActivityInfoRsp, func() any { return new(proto.GetActivityInfoRsp) })
	c.regMsg(ActivityPlayOpenAnimNotify, func() any { return new(proto.ActivityPlayOpenAnimNotify) })
	c.regMsg(ActivityInfoNotify, func() any { return new(proto.ActivityInfoNotify) })
	c.regMsg(ActivityScheduleInfoNotify, func() any { return new(proto.ActivityScheduleInfoNotify) })
	c.regMsg(ActivityTakeWatcherRewardReq, func() any { return new(proto.ActivityTakeWatcherRewardReq) })
	c.regMsg(ActivityTakeWatcherRewardRsp, func() any { return new(proto.ActivityTakeWatcherRewardRsp) })
	c.regMsg(ActivityUpdateWatcherNotify, func() any { return new(proto.ActivityUpdateWatcherNotify) })
	c.regMsg(SeaLampFlyLampReq, func() any { return new(proto.SeaLampFlyLampReq) })
	c.regMsg(SeaLampFlyLampRsp, func() any { return new(proto.SeaLampFlyLampRsp) })
	c.regMsg(SeaLampTakeContributionRewardReq, func() any { return new(proto.SeaLampTakeContributionRewardReq) })
	c.regMsg(SeaLampTakeContributionRewardRsp, func() any { return new(proto.SeaLampTakeContributionRewardRsp) })
	c.regMsg(SeaLampTakePhaseRewardReq, func() any { return new(proto.SeaLampTakePhaseRewardReq) })
	c.regMsg(SeaLampTakePhaseRewardRsp, func() any { return new(proto.SeaLampTakePhaseRewardRsp) })
	c.regMsg(SeaLampContributeItemReq, func() any { return new(proto.SeaLampContributeItemReq) })
	c.regMsg(SeaLampContributeItemRsp, func() any { return new(proto.SeaLampContributeItemRsp) })
	c.regMsg(LoadActivityTerrainNotify, func() any { return new(proto.LoadActivityTerrainNotify) })
	c.regMsg(ServerAnnounceNotify, func() any { return new(proto.ServerAnnounceNotify) })
	c.regMsg(ServerAnnounceRevokeNotify, func() any { return new(proto.ServerAnnounceRevokeNotify) })
	c.regMsg(SalesmanDeliverItemReq, func() any { return new(proto.SalesmanDeliverItemReq) })
	c.regMsg(SalesmanDeliverItemRsp, func() any { return new(proto.SalesmanDeliverItemRsp) })
	c.regMsg(SalesmanTakeRewardReq, func() any { return new(proto.SalesmanTakeRewardReq) })
	c.regMsg(SalesmanTakeRewardRsp, func() any { return new(proto.SalesmanTakeRewardRsp) })
	c.regMsg(ActivityCondStateChangeNotify, func() any { return new(proto.ActivityCondStateChangeNotify) })
	c.regMsg(EnterTrialAvatarActivityDungeonReq, func() any { return new(proto.EnterTrialAvatarActivityDungeonReq) })
	c.regMsg(EnterTrialAvatarActivityDungeonRsp, func() any { return new(proto.EnterTrialAvatarActivityDungeonRsp) })
	c.regMsg(ReceivedTrialAvatarActivityRewardReq, func() any { return new(proto.ReceivedTrialAvatarActivityRewardReq) })
	c.regMsg(ReceivedTrialAvatarActivityRewardRsp, func() any { return new(proto.ReceivedTrialAvatarActivityRewardRsp) })
	c.regMsg(TrialAvatarFirstPassDungeonNotify, func() any { return new(proto.TrialAvatarFirstPassDungeonNotify) })
	c.regMsg(TrialAvatarInDungeonIndexNotify, func() any { return new(proto.TrialAvatarInDungeonIndexNotify) })
	c.regMsg(AvatarAddNotify, func() any { return new(proto.AvatarAddNotify) })
	c.regMsg(AvatarDelNotify, func() any { return new(proto.AvatarDelNotify) })
	c.regMsg(SetUpAvatarTeamReq, func() any { return new(proto.SetUpAvatarTeamReq) })
	c.regMsg(SetUpAvatarTeamRsp, func() any { return new(proto.SetUpAvatarTeamRsp) })
	c.regMsg(ChooseCurAvatarTeamReq, func() any { return new(proto.ChooseCurAvatarTeamReq) })
	c.regMsg(ChooseCurAvatarTeamRsp, func() any { return new(proto.ChooseCurAvatarTeamRsp) })
	c.regMsg(ChangeAvatarReq, func() any { return new(proto.ChangeAvatarReq) })
	c.regMsg(ChangeAvatarRsp, func() any { return new(proto.ChangeAvatarRsp) })
	c.regMsg(AvatarPromoteReq, func() any { return new(proto.AvatarPromoteReq) })
	c.regMsg(AvatarPromoteRsp, func() any { return new(proto.AvatarPromoteRsp) })
	c.regMsg(SpringUseReq, func() any { return new(proto.SpringUseReq) })
	c.regMsg(SpringUseRsp, func() any { return new(proto.SpringUseRsp) })
	c.regMsg(RefreshBackgroundAvatarReq, func() any { return new(proto.RefreshBackgroundAvatarReq) })
	c.regMsg(RefreshBackgroundAvatarRsp, func() any { return new(proto.RefreshBackgroundAvatarRsp) })
	c.regMsg(AvatarTeamUpdateNotify, func() any { return new(proto.AvatarTeamUpdateNotify) })
	c.regMsg(AvatarDataNotify, func() any { return new(proto.AvatarDataNotify) })
	c.regMsg(AvatarUpgradeReq, func() any { return new(proto.AvatarUpgradeReq) })
	c.regMsg(AvatarUpgradeRsp, func() any { return new(proto.AvatarUpgradeRsp) })
	c.regMsg(AvatarDieAnimationEndReq, func() any { return new(proto.AvatarDieAnimationEndReq) })
	c.regMsg(AvatarDieAnimationEndRsp, func() any { return new(proto.AvatarDieAnimationEndRsp) })
	c.regMsg(AvatarChangeElementTypeReq, func() any { return new(proto.AvatarChangeElementTypeReq) })
	c.regMsg(AvatarChangeElementTypeRsp, func() any { return new(proto.AvatarChangeElementTypeRsp) })
	c.regMsg(AvatarFetterDataNotify, func() any { return new(proto.AvatarFetterDataNotify) })
	c.regMsg(AvatarExpeditionDataNotify, func() any { return new(proto.AvatarExpeditionDataNotify) })
	c.regMsg(AvatarExpeditionAllDataReq, func() any { return new(proto.AvatarExpeditionAllDataReq) })
	c.regMsg(AvatarExpeditionAllDataRsp, func() any { return new(proto.AvatarExpeditionAllDataRsp) })
	c.regMsg(AvatarExpeditionStartReq, func() any { return new(proto.AvatarExpeditionStartReq) })
	c.regMsg(AvatarExpeditionStartRsp, func() any { return new(proto.AvatarExpeditionStartRsp) })
	c.regMsg(AvatarExpeditionCallBackReq, func() any { return new(proto.AvatarExpeditionCallBackReq) })
	c.regMsg(AvatarExpeditionCallBackRsp, func() any { return new(proto.AvatarExpeditionCallBackRsp) })
	c.regMsg(AvatarExpeditionGetRewardReq, func() any { return new(proto.AvatarExpeditionGetRewardReq) })
	c.regMsg(AvatarExpeditionGetRewardRsp, func() any { return new(proto.AvatarExpeditionGetRewardRsp) })
	c.regMsg(ChangeMpTeamAvatarReq, func() any { return new(proto.ChangeMpTeamAvatarReq) })
	c.regMsg(ChangeMpTeamAvatarRsp, func() any { return new(proto.ChangeMpTeamAvatarRsp) })
	c.regMsg(ChangeTeamNameReq, func() any { return new(proto.ChangeTeamNameReq) })
	c.regMsg(ChangeTeamNameRsp, func() any { return new(proto.ChangeTeamNameRsp) })
	c.regMsg(SceneTeamUpdateNotify, func() any { return new(proto.SceneTeamUpdateNotify) })
	c.regMsg(FocusAvatarReq, func() any { return new(proto.FocusAvatarReq) })
	c.regMsg(FocusAvatarRsp, func() any { return new(proto.FocusAvatarRsp) })
	c.regMsg(AvatarSatiationDataNotify, func() any { return new(proto.AvatarSatiationDataNotify) })
	c.regMsg(AvatarWearFlycloakReq, func() any { return new(proto.AvatarWearFlycloakReq) })
	c.regMsg(AvatarWearFlycloakRsp, func() any { return new(proto.AvatarWearFlycloakRsp) })
	c.regMsg(AvatarFlycloakChangeNotify, func() any { return new(proto.AvatarFlycloakChangeNotify) })
	c.regMsg(AvatarGainFlycloakNotify, func() any { return new(proto.AvatarGainFlycloakNotify) })
	c.regMsg(AvatarEquipAffixStartNotify, func() any { return new(proto.AvatarEquipAffixStartNotify) })
	c.regMsg(AvatarFetterLevelRewardReq, func() any { return new(proto.AvatarFetterLevelRewardReq) })
	c.regMsg(AvatarFetterLevelRewardRsp, func() any { return new(proto.AvatarFetterLevelRewardRsp) })
	c.regMsg(BattlePassAllDataNotify, func() any { return new(proto.BattlePassAllDataNotify) })
	c.regMsg(BattlePassMissionUpdateNotify, func() any { return new(proto.BattlePassMissionUpdateNotify) })
	c.regMsg(BattlePassMissionDelNotify, func() any { return new(proto.BattlePassMissionDelNotify) })
	c.regMsg(BattlePassCurScheduleUpdateNotify, func() any { return new(proto.BattlePassCurScheduleUpdateNotify) })
	c.regMsg(TakeBattlePassRewardReq, func() any { return new(proto.TakeBattlePassRewardReq) })
	c.regMsg(TakeBattlePassRewardRsp, func() any { return new(proto.TakeBattlePassRewardRsp) })
	c.regMsg(TakeBattlePassMissionPointReq, func() any { return new(proto.TakeBattlePassMissionPointReq) })
	c.regMsg(TakeBattlePassMissionPointRsp, func() any { return new(proto.TakeBattlePassMissionPointRsp) })
	c.regMsg(GetBattlePassProductReq, func() any { return new(proto.GetBattlePassProductReq) })
	c.regMsg(GetBattlePassProductRsp, func() any { return new(proto.GetBattlePassProductRsp) })
	c.regMsg(SetBattlePassViewedReq, func() any { return new(proto.SetBattlePassViewedReq) })
	c.regMsg(SetBattlePassViewedRsp, func() any { return new(proto.SetBattlePassViewedRsp) })
	c.regMsg(BattlePassBuySuccNotify, func() any { return new(proto.BattlePassBuySuccNotify) })
	c.regMsg(BuyBattlePassLevelReq, func() any { return new(proto.BuyBattlePassLevelReq) })
	c.regMsg(BuyBattlePassLevelRsp, func() any { return new(proto.BuyBattlePassLevelRsp) })
	c.regMsg(TowerBriefDataNotify, func() any { return new(proto.TowerBriefDataNotify) })
	c.regMsg(TowerFloorRecordChangeNotify, func() any { return new(proto.TowerFloorRecordChangeNotify) })
	c.regMsg(TowerCurLevelRecordChangeNotify, func() any { return new(proto.TowerCurLevelRecordChangeNotify) })
	c.regMsg(TowerDailyRewardProgressChangeNotify, func() any { return new(proto.TowerDailyRewardProgressChangeNotify) })
	c.regMsg(TowerTeamSelectReq, func() any { return new(proto.TowerTeamSelectReq) })
	c.regMsg(TowerTeamSelectRsp, func() any { return new(proto.TowerTeamSelectRsp) })
	c.regMsg(TowerAllDataReq, func() any { return new(proto.TowerAllDataReq) })
	c.regMsg(TowerAllDataRsp, func() any { return new(proto.TowerAllDataRsp) })
	c.regMsg(TowerEnterLevelReq, func() any { return new(proto.TowerEnterLevelReq) })
	c.regMsg(TowerEnterLevelRsp, func() any { return new(proto.TowerEnterLevelRsp) })
	c.regMsg(TowerBuffSelectReq, func() any { return new(proto.TowerBuffSelectReq) })
	c.regMsg(TowerBuffSelectRsp, func() any { return new(proto.TowerBuffSelectRsp) })
	c.regMsg(TowerSurrenderReq, func() any { return new(proto.TowerSurrenderReq) })
	c.regMsg(TowerSurrenderRsp, func() any { return new(proto.TowerSurrenderRsp) })
	c.regMsg(TowerGetFloorStarRewardReq, func() any { return new(proto.TowerGetFloorStarRewardReq) })
	c.regMsg(TowerGetFloorStarRewardRsp, func() any { return new(proto.TowerGetFloorStarRewardRsp) })
	c.regMsg(TowerLevelEndNotify, func() any { return new(proto.TowerLevelEndNotify) })
	c.regMsg(TowerLevelStarCondNotify, func() any { return new(proto.TowerLevelStarCondNotify) })
	c.regMsg(TowerMiddleLevelChangeTeamNotify, func() any { return new(proto.TowerMiddleLevelChangeTeamNotify) })
	c.regMsg(TowerRecordHandbookReq, func() any { return new(proto.TowerRecordHandbookReq) })
	c.regMsg(TowerRecordHandbookRsp, func() any { return new(proto.TowerRecordHandbookRsp) })
	c.regMsg(WatcherAllDataNotify, func() any { return new(proto.WatcherAllDataNotify) })
	c.regMsg(WatcherChangeNotify, func() any { return new(proto.WatcherChangeNotify) })
	c.regMsg(WatcherEventNotify, func() any { return new(proto.WatcherEventNotify) })
	c.regMsg(WatcherEventTypeNotify, func() any { return new(proto.WatcherEventTypeNotify) })
	c.regMsg(PushTipsAllDataNotify, func() any { return new(proto.PushTipsAllDataNotify) })
	c.regMsg(PushTipsChangeNotify, func() any { return new(proto.PushTipsChangeNotify) })
	c.regMsg(PushTipsReadFinishReq, func() any { return new(proto.PushTipsReadFinishReq) })
	c.regMsg(PushTipsReadFinishRsp, func() any { return new(proto.PushTipsReadFinishRsp) })
	c.regMsg(GetPushTipsRewardReq, func() any { return new(proto.GetPushTipsRewardReq) })
	c.regMsg(GetPushTipsRewardRsp, func() any { return new(proto.GetPushTipsRewardRsp) })
	c.regMsg(GetBlossomBriefInfoListReq, func() any { return new(proto.GetBlossomBriefInfoListReq) })
	c.regMsg(GetBlossomBriefInfoListRsp, func() any { return new(proto.GetBlossomBriefInfoListRsp) })
	c.regMsg(BlossomBriefInfoNotify, func() any { return new(proto.BlossomBriefInfoNotify) })
	c.regMsg(WorldOwnerBlossomBriefInfoNotify, func() any { return new(proto.WorldOwnerBlossomBriefInfoNotify) })
	c.regMsg(WorldOwnerBlossomScheduleInfoNotify, func() any { return new(proto.WorldOwnerBlossomScheduleInfoNotify) })
	c.regMsg(BlossomChestCreateNotify, func() any { return new(proto.BlossomChestCreateNotify) })
	c.regMsg(OpenBlossomCircleCampGuideNotify, func() any { return new(proto.OpenBlossomCircleCampGuideNotify) })
	c.regMsg(CodexDataFullNotify, func() any { return new(proto.CodexDataFullNotify) })
	c.regMsg(CodexDataUpdataNotify, func() any { return new(proto.CodexDataUpdataNotify) })
	c.regMsg(QueryCodexMonsterBeKilledNumReq, func() any { return new(proto.QueryCodexMonsterBeKilledNumReq) })
	c.regMsg(QueryCodexMonsterBeKilledNumRsp, func() any { return new(proto.QueryCodexMonsterBeKilledNumRsp) })
	c.regMsg(DungeonEntryInfoReq, func() any { return new(proto.DungeonEntryInfoReq) })
	c.regMsg(DungeonEntryInfoRsp, func() any { return new(proto.DungeonEntryInfoRsp) })
	c.regMsg(PlayerEnterDungeonReq, func() any { return new(proto.PlayerEnterDungeonReq) })
	c.regMsg(PlayerEnterDungeonRsp, func() any { return new(proto.PlayerEnterDungeonRsp) })
	c.regMsg(PlayerQuitDungeonReq, func() any { return new(proto.PlayerQuitDungeonReq) })
	c.regMsg(PlayerQuitDungeonRsp, func() any { return new(proto.PlayerQuitDungeonRsp) })
	c.regMsg(DungeonWayPointNotify, func() any { return new(proto.DungeonWayPointNotify) })
	c.regMsg(DungeonWayPointActivateReq, func() any { return new(proto.DungeonWayPointActivateReq) })
	c.regMsg(DungeonWayPointActivateRsp, func() any { return new(proto.DungeonWayPointActivateRsp) })
	c.regMsg(DungeonSettleNotify, func() any { return new(proto.DungeonSettleNotify) })
	c.regMsg(DungeonPlayerDieNotify, func() any { return new(proto.DungeonPlayerDieNotify) })
	c.regMsg(DungeonDieOptionReq, func() any { return new(proto.DungeonDieOptionReq) })
	c.regMsg(DungeonDieOptionRsp, func() any { return new(proto.DungeonDieOptionRsp) })
	c.regMsg(DungeonShowReminderNotify, func() any { return new(proto.DungeonShowReminderNotify) })
	c.regMsg(DungeonPlayerDieReq, func() any { return new(proto.DungeonPlayerDieReq) })
	c.regMsg(DungeonPlayerDieRsp, func() any { return new(proto.DungeonPlayerDieRsp) })
	c.regMsg(DungeonDataNotify, func() any { return new(proto.DungeonDataNotify) })
	c.regMsg(DungeonChallengeBeginNotify, func() any { return new(proto.DungeonChallengeBeginNotify) })
	c.regMsg(DungeonChallengeFinishNotify, func() any { return new(proto.DungeonChallengeFinishNotify) })
	c.regMsg(ChallengeDataNotify, func() any { return new(proto.ChallengeDataNotify) })
	c.regMsg(DungeonFollowNotify, func() any { return new(proto.DungeonFollowNotify) })
	c.regMsg(DungeonGetStatueDropReq, func() any { return new(proto.DungeonGetStatueDropReq) })
	c.regMsg(DungeonGetStatueDropRsp, func() any { return new(proto.DungeonGetStatueDropRsp) })
	c.regMsg(ChallengeRecordNotify, func() any { return new(proto.ChallengeRecordNotify) })
	c.regMsg(DungeonCandidateTeamInfoNotify, func() any { return new(proto.DungeonCandidateTeamInfoNotify) })
	c.regMsg(DungeonCandidateTeamInviteNotify, func() any { return new(proto.DungeonCandidateTeamInviteNotify) })
	c.regMsg(DungeonCandidateTeamRefuseNotify, func() any { return new(proto.DungeonCandidateTeamRefuseNotify) })
	c.regMsg(DungeonCandidateTeamPlayerLeaveNotify, func() any { return new(proto.DungeonCandidateTeamPlayerLeaveNotify) })
	c.regMsg(DungeonCandidateTeamDismissNotify, func() any { return new(proto.DungeonCandidateTeamDismissNotify) })
	c.regMsg(DungeonCandidateTeamCreateReq, func() any { return new(proto.DungeonCandidateTeamCreateReq) })
	c.regMsg(DungeonCandidateTeamCreateRsp, func() any { return new(proto.DungeonCandidateTeamCreateRsp) })
	c.regMsg(DungeonCandidateTeamInviteReq, func() any { return new(proto.DungeonCandidateTeamInviteReq) })
	c.regMsg(DungeonCandidateTeamInviteRsp, func() any { return new(proto.DungeonCandidateTeamInviteRsp) })
	c.regMsg(DungeonCandidateTeamKickReq, func() any { return new(proto.DungeonCandidateTeamKickReq) })
	c.regMsg(DungeonCandidateTeamKickRsp, func() any { return new(proto.DungeonCandidateTeamKickRsp) })
	c.regMsg(DungeonCandidateTeamLeaveReq, func() any { return new(proto.DungeonCandidateTeamLeaveReq) })
	c.regMsg(DungeonCandidateTeamLeaveRsp, func() any { return new(proto.DungeonCandidateTeamLeaveRsp) })
	c.regMsg(DungeonCandidateTeamReplyInviteReq, func() any { return new(proto.DungeonCandidateTeamReplyInviteReq) })
	c.regMsg(DungeonCandidateTeamReplyInviteRsp, func() any { return new(proto.DungeonCandidateTeamReplyInviteRsp) })
	c.regMsg(DungeonCandidateTeamSetReadyReq, func() any { return new(proto.DungeonCandidateTeamSetReadyReq) })
	c.regMsg(DungeonCandidateTeamSetReadyRsp, func() any { return new(proto.DungeonCandidateTeamSetReadyRsp) })
	c.regMsg(DungeonCandidateTeamChangeAvatarReq, func() any { return new(proto.DungeonCandidateTeamChangeAvatarReq) })
	c.regMsg(DungeonCandidateTeamChangeAvatarRsp, func() any { return new(proto.DungeonCandidateTeamChangeAvatarRsp) })
	c.regMsg(GetDailyDungeonEntryInfoReq, func() any { return new(proto.GetDailyDungeonEntryInfoReq) })
	c.regMsg(GetDailyDungeonEntryInfoRsp, func() any { return new(proto.GetDailyDungeonEntryInfoRsp) })
	c.regMsg(DungeonSlipRevivePointActivateReq, func() any { return new(proto.DungeonSlipRevivePointActivateReq) })
	c.regMsg(DungeonSlipRevivePointActivateRsp, func() any { return new(proto.DungeonSlipRevivePointActivateRsp) })
	c.regMsg(DungeonInterruptChallengeReq, func() any { return new(proto.DungeonInterruptChallengeReq) })
	c.regMsg(DungeonInterruptChallengeRsp, func() any { return new(proto.DungeonInterruptChallengeRsp) })
	c.regMsg(InteractDailyDungeonInfoNotify, func() any { return new(proto.InteractDailyDungeonInfoNotify) })
	c.regMsg(EvtBeingHitNotify, func() any { return new(proto.EvtBeingHitNotify) })
	c.regMsg(EvtAnimatorParameterNotify, func() any { return new(proto.EvtAnimatorParameterNotify) })
	c.regMsg(HostPlayerNotify, func() any { return new(proto.HostPlayerNotify) })
	c.regMsg(EvtDoSkillSuccNotify, func() any { return new(proto.EvtDoSkillSuccNotify) })
	c.regMsg(EvtCreateGadgetNotify, func() any { return new(proto.EvtCreateGadgetNotify) })
	c.regMsg(EvtDestroyGadgetNotify, func() any { return new(proto.EvtDestroyGadgetNotify) })
	c.regMsg(EvtFaceToEntityNotify, func() any { return new(proto.EvtFaceToEntityNotify) })
	c.regMsg(EvtFaceToDirNotify, func() any { return new(proto.EvtFaceToDirNotify) })
	c.regMsg(EvtCostStaminaNotify, func() any { return new(proto.EvtCostStaminaNotify) })
	c.regMsg(EvtSetAttackTargetNotify, func() any { return new(proto.EvtSetAttackTargetNotify) })
	c.regMsg(EvtAnimatorStateChangedNotify, func() any { return new(proto.EvtAnimatorStateChangedNotify) })
	c.regMsg(EvtRushMoveNotify, func() any { return new(proto.EvtRushMoveNotify) })
	c.regMsg(EvtBulletHitNotify, func() any { return new(proto.EvtBulletHitNotify) })
	c.regMsg(EvtBulletDeactiveNotify, func() any { return new(proto.EvtBulletDeactiveNotify) })
	c.regMsg(EvtEntityStartDieEndNotify, func() any { return new(proto.EvtEntityStartDieEndNotify) })
	c.regMsg(EvtBulletMoveNotify, func() any { return new(proto.EvtBulletMoveNotify) })
	c.regMsg(EvtAvatarEnterFocusNotify, func() any { return new(proto.EvtAvatarEnterFocusNotify) })
	c.regMsg(EvtAvatarExitFocusNotify, func() any { return new(proto.EvtAvatarExitFocusNotify) })
	c.regMsg(EvtAvatarUpdateFocusNotify, func() any { return new(proto.EvtAvatarUpdateFocusNotify) })
	c.regMsg(EntityAuthorityChangeNotify, func() any { return new(proto.EntityAuthorityChangeNotify) })
	c.regMsg(AvatarBuffAddNotify, func() any { return new(proto.AvatarBuffAddNotify) })
	c.regMsg(AvatarBuffDelNotify, func() any { return new(proto.AvatarBuffDelNotify) })
	c.regMsg(MonsterAlertChangeNotify, func() any { return new(proto.MonsterAlertChangeNotify) })
	c.regMsg(MonsterForceAlertNotify, func() any { return new(proto.MonsterForceAlertNotify) })
	c.regMsg(AvatarEnterElementViewNotify, func() any { return new(proto.AvatarEnterElementViewNotify) })
	c.regMsg(TriggerCreateGadgetToEquipPartNotify, func() any { return new(proto.TriggerCreateGadgetToEquipPartNotify) })
	c.regMsg(EvtEntityRenderersChangedNotify, func() any { return new(proto.EvtEntityRenderersChangedNotify) })
	c.regMsg(AnimatorForceSetAirMoveNotify, func() any { return new(proto.AnimatorForceSetAirMoveNotify) })
	c.regMsg(EvtAiSyncSkillCdNotify, func() any { return new(proto.EvtAiSyncSkillCdNotify) })
	c.regMsg(EvtBeingHitsCombineNotify, func() any { return new(proto.EvtBeingHitsCombineNotify) })
	c.regMsg(EvtAvatarSitDownNotify, func() any { return new(proto.EvtAvatarSitDownNotify) })
	c.regMsg(EvtAvatarStandUpNotify, func() any { return new(proto.EvtAvatarStandUpNotify) })
	c.regMsg(CreateMassiveEntityReq, func() any { return new(proto.CreateMassiveEntityReq) })
	c.regMsg(CreateMassiveEntityRsp, func() any { return new(proto.CreateMassiveEntityRsp) })
	c.regMsg(CreateMassiveEntityNotify, func() any { return new(proto.CreateMassiveEntityNotify) })
	c.regMsg(DestroyMassiveEntityNotify, func() any { return new(proto.DestroyMassiveEntityNotify) })
	c.regMsg(MassiveEntityStateChangedNotify, func() any { return new(proto.MassiveEntityStateChangedNotify) })
	c.regMsg(SyncTeamEntityNotify, func() any { return new(proto.SyncTeamEntityNotify) })
	c.regMsg(DelTeamEntityNotify, func() any { return new(proto.DelTeamEntityNotify) })
	c.regMsg(CombatInvocationsNotify, func() any { return new(proto.CombatInvocationsNotify) })
	c.regMsg(ServerBuffChangeNotify, func() any { return new(proto.ServerBuffChangeNotify) })
	c.regMsg(EvtAiSyncCombatThreatInfoNotify, func() any { return new(proto.EvtAiSyncCombatThreatInfoNotify) })
	c.regMsg(MassiveEntityElementOpBatchNotify, func() any { return new(proto.MassiveEntityElementOpBatchNotify) })
	c.regMsg(EntityAiSyncNotify, func() any { return new(proto.EntityAiSyncNotify) })
	c.regMsg(GetGachaInfoReq, func() any { return new(proto.GetGachaInfoReq) })
	c.regMsg(GetGachaInfoRsp, func() any { return new(proto.GetGachaInfoRsp) })
	c.regMsg(DoGachaReq, func() any { return new(proto.DoGachaReq) })
	c.regMsg(DoGachaRsp, func() any { return new(proto.DoGachaRsp) })
	c.regMsg(GadgetInteractReq, func() any { return new(proto.GadgetInteractReq) })
	c.regMsg(GadgetInteractRsp, func() any { return new(proto.GadgetInteractRsp) })
	c.regMsg(GadgetStateNotify, func() any { return new(proto.GadgetStateNotify) })
	c.regMsg(WorktopOptionNotify, func() any { return new(proto.WorktopOptionNotify) })
	c.regMsg(SelectWorktopOptionReq, func() any { return new(proto.SelectWorktopOptionReq) })
	c.regMsg(SelectWorktopOptionRsp, func() any { return new(proto.SelectWorktopOptionRsp) })
	c.regMsg(BossChestActivateNotify, func() any { return new(proto.BossChestActivateNotify) })
	c.regMsg(BlossomChestInfoNotify, func() any { return new(proto.BlossomChestInfoNotify) })
	c.regMsg(GadgetPlayStartNotify, func() any { return new(proto.GadgetPlayStartNotify) })
	c.regMsg(GadgetPlayStopNotify, func() any { return new(proto.GadgetPlayStopNotify) })
	c.regMsg(GadgetPlayDataNotify, func() any { return new(proto.GadgetPlayDataNotify) })
	c.regMsg(GadgetPlayUidOpNotify, func() any { return new(proto.GadgetPlayUidOpNotify) })
	c.regMsg(PlayerInvestigationAllInfoNotify, func() any { return new(proto.PlayerInvestigationAllInfoNotify) })
	c.regMsg(TakeInvestigationRewardReq, func() any { return new(proto.TakeInvestigationRewardReq) })
	c.regMsg(TakeInvestigationRewardRsp, func() any { return new(proto.TakeInvestigationRewardRsp) })
	c.regMsg(TakeInvestigationTargetRewardReq, func() any { return new(proto.TakeInvestigationTargetRewardReq) })
	c.regMsg(TakeInvestigationTargetRewardRsp, func() any { return new(proto.TakeInvestigationTargetRewardRsp) })
	c.regMsg(GetInvestigationMonsterReq, func() any { return new(proto.GetInvestigationMonsterReq) })
	c.regMsg(GetInvestigationMonsterRsp, func() any { return new(proto.GetInvestigationMonsterRsp) })
	c.regMsg(PlayerInvestigationNotify, func() any { return new(proto.PlayerInvestigationNotify) })
	c.regMsg(PlayerInvestigationTargetNotify, func() any { return new(proto.PlayerInvestigationTargetNotify) })
	c.regMsg(PlayerStoreNotify, func() any { return new(proto.PlayerStoreNotify) })
	c.regMsg(StoreWeightLimitNotify, func() any { return new(proto.StoreWeightLimitNotify) })
	c.regMsg(StoreItemChangeNotify, func() any { return new(proto.StoreItemChangeNotify) })
	c.regMsg(StoreItemDelNotify, func() any { return new(proto.StoreItemDelNotify) })
	c.regMsg(ItemAddHintNotify, func() any { return new(proto.ItemAddHintNotify) })
	c.regMsg(UseItemReq, func() any { return new(proto.UseItemReq) })
	c.regMsg(UseItemRsp, func() any { return new(proto.UseItemRsp) })
	c.regMsg(DropItemReq, func() any { return new(proto.DropItemReq) })
	c.regMsg(DropItemRsp, func() any { return new(proto.DropItemRsp) })
	c.regMsg(WearEquipReq, func() any { return new(proto.WearEquipReq) })
	c.regMsg(WearEquipRsp, func() any { return new(proto.WearEquipRsp) })
	c.regMsg(TakeoffEquipReq, func() any { return new(proto.TakeoffEquipReq) })
	c.regMsg(TakeoffEquipRsp, func() any { return new(proto.TakeoffEquipRsp) })
	c.regMsg(AvatarEquipChangeNotify, func() any { return new(proto.AvatarEquipChangeNotify) })
	c.regMsg(WeaponUpgradeReq, func() any { return new(proto.WeaponUpgradeReq) })
	c.regMsg(WeaponUpgradeRsp, func() any { return new(proto.WeaponUpgradeRsp) })
	c.regMsg(WeaponPromoteReq, func() any { return new(proto.WeaponPromoteReq) })
	c.regMsg(WeaponPromoteRsp, func() any { return new(proto.WeaponPromoteRsp) })
	c.regMsg(ReliquaryUpgradeReq, func() any { return new(proto.ReliquaryUpgradeReq) })
	c.regMsg(ReliquaryUpgradeRsp, func() any { return new(proto.ReliquaryUpgradeRsp) })
	c.regMsg(ReliquaryPromoteReq, func() any { return new(proto.ReliquaryPromoteReq) })
	c.regMsg(ReliquaryPromoteRsp, func() any { return new(proto.ReliquaryPromoteRsp) })
	c.regMsg(AvatarCardChangeReq, func() any { return new(proto.AvatarCardChangeReq) })
	c.regMsg(AvatarCardChangeRsp, func() any { return new(proto.AvatarCardChangeRsp) })
	c.regMsg(GrantRewardNotify, func() any { return new(proto.GrantRewardNotify) })
	c.regMsg(WeaponAwakenReq, func() any { return new(proto.WeaponAwakenReq) })
	c.regMsg(WeaponAwakenRsp, func() any { return new(proto.WeaponAwakenRsp) })
	c.regMsg(ItemCdGroupTimeNotify, func() any { return new(proto.ItemCdGroupTimeNotify) })
	c.regMsg(DropHintNotify, func() any { return new(proto.DropHintNotify) })
	c.regMsg(CombineReq, func() any { return new(proto.CombineReq) })
	c.regMsg(CombineRsp, func() any { return new(proto.CombineRsp) })
	c.regMsg(ForgeQueueDataNotify, func() any { return new(proto.ForgeQueueDataNotify) })
	c.regMsg(ForgeGetQueueDataReq, func() any { return new(proto.ForgeGetQueueDataReq) })
	c.regMsg(ForgeGetQueueDataRsp, func() any { return new(proto.ForgeGetQueueDataRsp) })
	c.regMsg(ForgeStartReq, func() any { return new(proto.ForgeStartReq) })
	c.regMsg(ForgeStartRsp, func() any { return new(proto.ForgeStartRsp) })
	c.regMsg(ForgeQueueManipulateReq, func() any { return new(proto.ForgeQueueManipulateReq) })
	c.regMsg(ForgeQueueManipulateRsp, func() any { return new(proto.ForgeQueueManipulateRsp) })
	c.regMsg(ResinChangeNotify, func() any { return new(proto.ResinChangeNotify) })
	c.regMsg(BuyResinReq, func() any { return new(proto.BuyResinReq) })
	c.regMsg(BuyResinRsp, func() any { return new(proto.BuyResinRsp) })
	c.regMsg(MaterialDeleteReturnNotify, func() any { return new(proto.MaterialDeleteReturnNotify) })
	c.regMsg(TakeMaterialDeleteReturnReq, func() any { return new(proto.TakeMaterialDeleteReturnReq) })
	c.regMsg(TakeMaterialDeleteReturnRsp, func() any { return new(proto.TakeMaterialDeleteReturnRsp) })
	c.regMsg(MaterialDeleteUpdateNotify, func() any { return new(proto.MaterialDeleteUpdateNotify) })
	c.regMsg(McoinExchangeHcoinReq, func() any { return new(proto.McoinExchangeHcoinReq) })
	c.regMsg(McoinExchangeHcoinRsp, func() any { return new(proto.McoinExchangeHcoinRsp) })
	c.regMsg(DestroyMaterialReq, func() any { return new(proto.DestroyMaterialReq) })
	c.regMsg(DestroyMaterialRsp, func() any { return new(proto.DestroyMaterialRsp) })
	c.regMsg(MailChangeNotify, func() any { return new(proto.MailChangeNotify) })
	c.regMsg(ReadMailNotify, func() any { return new(proto.ReadMailNotify) })
	c.regMsg(GetMailItemReq, func() any { return new(proto.GetMailItemReq) })
	c.regMsg(GetMailItemRsp, func() any { return new(proto.GetMailItemRsp) })
	c.regMsg(DelMailReq, func() any { return new(proto.DelMailReq) })
	c.regMsg(DelMailRsp, func() any { return new(proto.DelMailRsp) })
	c.regMsg(GetAuthkeyReq, func() any { return new(proto.GetAuthkeyReq) })
	c.regMsg(GetAuthkeyRsp, func() any { return new(proto.GetAuthkeyRsp) })
	c.regMsg(ClientNewMailNotify, func() any { return new(proto.ClientNewMailNotify) })
	c.regMsg(GetAllMailReq, func() any { return new(proto.GetAllMailReq) })
	c.regMsg(GetAllMailRsp, func() any { return new(proto.GetAllMailRsp) })
	c.regMsg(PlayerStartMatchReq, func() any { return new(proto.PlayerStartMatchReq) })
	c.regMsg(PlayerStartMatchRsp, func() any { return new(proto.PlayerStartMatchRsp) })
	c.regMsg(PlayerMatchInfoNotify, func() any { return new(proto.PlayerMatchInfoNotify) })
	c.regMsg(PlayerCancelMatchReq, func() any { return new(proto.PlayerCancelMatchReq) })
	c.regMsg(PlayerCancelMatchRsp, func() any { return new(proto.PlayerCancelMatchRsp) })
	c.regMsg(PlayerMatchStopNotify, func() any { return new(proto.PlayerMatchStopNotify) })
	c.regMsg(PlayerMatchSuccNotify, func() any { return new(proto.PlayerMatchSuccNotify) })
	c.regMsg(PlayerConfirmMatchReq, func() any { return new(proto.PlayerConfirmMatchReq) })
	c.regMsg(PlayerConfirmMatchRsp, func() any { return new(proto.PlayerConfirmMatchRsp) })
	c.regMsg(PlayerAllowEnterMpAfterAgreeMatchNotify, func() any { return new(proto.PlayerAllowEnterMpAfterAgreeMatchNotify) })
	c.regMsg(PlayerMatchAgreedResultNotify, func() any { return new(proto.PlayerMatchAgreedResultNotify) })
	c.regMsg(PlayerApplyEnterMpAfterMatchAgreedNotify, func() any { return new(proto.PlayerApplyEnterMpAfterMatchAgreedNotify) })
	c.regMsg(KeepAliveNotify, func() any { return new(proto.KeepAliveNotify) })
	c.regMsg(GmTalkReq, func() any { return new(proto.GmTalkReq) })
	c.regMsg(GmTalkRsp, func() any { return new(proto.GmTalkRsp) })
	c.regMsg(ShowMessageNotify, func() any { return new(proto.ShowMessageNotify) })
	c.regMsg(PingReq, func() any { return new(proto.PingReq) })
	c.regMsg(PingRsp, func() any { return new(proto.PingRsp) })
	c.regMsg(GetOnlinePlayerListReq, func() any { return new(proto.GetOnlinePlayerListReq) })
	c.regMsg(GetOnlinePlayerListRsp, func() any { return new(proto.GetOnlinePlayerListRsp) })
	c.regMsg(ServerTimeNotify, func() any { return new(proto.ServerTimeNotify) })
	c.regMsg(ServerLogNotify, func() any { return new(proto.ServerLogNotify) })
	c.regMsg(ClientReconnectNotify, func() any { return new(proto.ClientReconnectNotify) })
	c.regMsg(RobotPushPlayerDataNotify, func() any { return new(proto.RobotPushPlayerDataNotify) })
	c.regMsg(ClientReportNotify, func() any { return new(proto.ClientReportNotify) })
	c.regMsg(UnionCmdNotify, func() any { return new(proto.UnionCmdNotify) })
	c.regMsg(GetOnlinePlayerInfoReq, func() any { return new(proto.GetOnlinePlayerInfoReq) })
	c.regMsg(GetOnlinePlayerInfoRsp, func() any { return new(proto.GetOnlinePlayerInfoRsp) })
	c.regMsg(CheckSegmentCRCNotify, func() any { return new(proto.CheckSegmentCRCNotify) })
	c.regMsg(CheckSegmentCRCReq, func() any { return new(proto.CheckSegmentCRCReq) })
	c.regMsg(WorldPlayerRTTNotify, func() any { return new(proto.WorldPlayerRTTNotify) })
	c.regMsg(EchoNotify, func() any { return new(proto.EchoNotify) })
	c.regMsg(MonsterSummonTagNotify, func() any { return new(proto.MonsterSummonTagNotify) })
	c.regMsg(PlayerApplyEnterMpNotify, func() any { return new(proto.PlayerApplyEnterMpNotify) })
	c.regMsg(PlayerApplyEnterMpReq, func() any { return new(proto.PlayerApplyEnterMpReq) })
	c.regMsg(PlayerApplyEnterMpRsp, func() any { return new(proto.PlayerApplyEnterMpRsp) })
	c.regMsg(PlayerApplyEnterMpResultNotify, func() any { return new(proto.PlayerApplyEnterMpResultNotify) })
	c.regMsg(PlayerApplyEnterMpResultReq, func() any { return new(proto.PlayerApplyEnterMpResultReq) })
	c.regMsg(PlayerApplyEnterMpResultRsp, func() any { return new(proto.PlayerApplyEnterMpResultRsp) })
	c.regMsg(PlayerQuitFromMpNotify, func() any { return new(proto.PlayerQuitFromMpNotify) })
	c.regMsg(PlayerPreEnterMpNotify, func() any { return new(proto.PlayerPreEnterMpNotify) })
	c.regMsg(GetPlayerMpModeAvailabilityReq, func() any { return new(proto.GetPlayerMpModeAvailabilityReq) })
	c.regMsg(GetPlayerMpModeAvailabilityRsp, func() any { return new(proto.GetPlayerMpModeAvailabilityRsp) })
	c.regMsg(PlayerSetOnlyMPWithPSPlayerReq, func() any { return new(proto.PlayerSetOnlyMPWithPSPlayerReq) })
	c.regMsg(PlayerSetOnlyMPWithPSPlayerRsp, func() any { return new(proto.PlayerSetOnlyMPWithPSPlayerRsp) })
	c.regMsg(PSPlayerApplyEnterMpReq, func() any { return new(proto.PSPlayerApplyEnterMpReq) })
	c.regMsg(PSPlayerApplyEnterMpRsp, func() any { return new(proto.PSPlayerApplyEnterMpRsp) })
	c.regMsg(MpPlayOwnerCheckReq, func() any { return new(proto.MpPlayOwnerCheckReq) })
	c.regMsg(MpPlayOwnerCheckRsp, func() any { return new(proto.MpPlayOwnerCheckRsp) })
	c.regMsg(MpPlayOwnerStartInviteReq, func() any { return new(proto.MpPlayOwnerStartInviteReq) })
	c.regMsg(MpPlayOwnerStartInviteRsp, func() any { return new(proto.MpPlayOwnerStartInviteRsp) })
	c.regMsg(MpPlayOwnerInviteNotify, func() any { return new(proto.MpPlayOwnerInviteNotify) })
	c.regMsg(MpPlayGuestReplyInviteReq, func() any { return new(proto.MpPlayGuestReplyInviteReq) })
	c.regMsg(MpPlayGuestReplyInviteRsp, func() any { return new(proto.MpPlayGuestReplyInviteRsp) })
	c.regMsg(MpPlayGuestReplyNotify, func() any { return new(proto.MpPlayGuestReplyNotify) })
	c.regMsg(MpPlayPrepareNotify, func() any { return new(proto.MpPlayPrepareNotify) })
	c.regMsg(MpPlayInviteResultNotify, func() any { return new(proto.MpPlayInviteResultNotify) })
	c.regMsg(MpPlayPrepareInterruptNotify, func() any { return new(proto.MpPlayPrepareInterruptNotify) })
	c.regMsg(NpcTalkReq, func() any { return new(proto.NpcTalkReq) })
	c.regMsg(NpcTalkRsp, func() any { return new(proto.NpcTalkRsp) })
	c.regMsg(GetSceneNpcPositionReq, func() any { return new(proto.GetSceneNpcPositionReq) })
	c.regMsg(GetSceneNpcPositionRsp, func() any { return new(proto.GetSceneNpcPositionRsp) })
	c.regMsg(QueryPathReq, func() any { return new(proto.QueryPathReq) })
	c.regMsg(QueryPathRsp, func() any { return new(proto.QueryPathRsp) })
	c.regMsg(ObstacleModifyNotify, func() any { return new(proto.ObstacleModifyNotify) })
	c.regMsg(PathfindingPingNotify, func() any { return new(proto.PathfindingPingNotify) })
	c.regMsg(PathfindingEnterSceneReq, func() any { return new(proto.PathfindingEnterSceneReq) })
	c.regMsg(PathfindingEnterSceneRsp, func() any { return new(proto.PathfindingEnterSceneRsp) })
	c.regMsg(GMShowObstacleReq, func() any { return new(proto.GMShowObstacleReq) })
	c.regMsg(GMShowObstacleRsp, func() any { return new(proto.GMShowObstacleRsp) })
	c.regMsg(GMShowNavMeshReq, func() any { return new(proto.GMShowNavMeshReq) })
	c.regMsg(GMShowNavMeshRsp, func() any { return new(proto.GMShowNavMeshRsp) })
	c.regMsg(NavMeshStatsNotify, func() any { return new(proto.NavMeshStatsNotify) })
	c.regMsg(GetPlayerTokenRsp, func() any { return new(proto.GetPlayerTokenRsp) })
	c.regMsg(PlayerLoginReq, func() any { return new(proto.PlayerLoginReq) })
	c.regMsg(PlayerLoginRsp, func() any { return new(proto.PlayerLoginRsp) })
	c.regMsg(PlayerLogoutReq, func() any { return new(proto.PlayerLogoutReq) })
	c.regMsg(PlayerLogoutRsp, func() any { return new(proto.PlayerLogoutRsp) })
	c.regMsg(PlayerLogoutNotify, func() any { return new(proto.PlayerLogoutNotify) })
	c.regMsg(PlayerDataNotify, func() any { return new(proto.PlayerDataNotify) })
	c.regMsg(ChangeGameTimeReq, func() any { return new(proto.ChangeGameTimeReq) })
	c.regMsg(ChangeGameTimeRsp, func() any { return new(proto.ChangeGameTimeRsp) })
	c.regMsg(PlayerGameTimeNotify, func() any { return new(proto.PlayerGameTimeNotify) })
	c.regMsg(PlayerPropNotify, func() any { return new(proto.PlayerPropNotify) })
	c.regMsg(ClientTriggerEventNotify, func() any { return new(proto.ClientTriggerEventNotify) })
	c.regMsg(SetPlayerPropReq, func() any { return new(proto.SetPlayerPropReq) })
	c.regMsg(SetPlayerPropRsp, func() any { return new(proto.SetPlayerPropRsp) })
	c.regMsg(SetPlayerBornDataReq, func() any { return new(proto.SetPlayerBornDataReq) })
	c.regMsg(SetPlayerBornDataRsp, func() any { return new(proto.SetPlayerBornDataRsp) })
	c.regMsg(DoSetPlayerBornDataNotify, func() any { return new(proto.DoSetPlayerBornDataNotify) })
	c.regMsg(PlayerPropChangeNotify, func() any { return new(proto.PlayerPropChangeNotify) })
	c.regMsg(SetPlayerNameReq, func() any { return new(proto.SetPlayerNameReq) })
	c.regMsg(SetPlayerNameRsp, func() any { return new(proto.SetPlayerNameRsp) })
	c.regMsg(SetOpenStateReq, func() any { return new(proto.SetOpenStateReq) })
	c.regMsg(SetOpenStateRsp, func() any { return new(proto.SetOpenStateRsp) })
	c.regMsg(OpenStateUpdateNotify, func() any { return new(proto.OpenStateUpdateNotify) })
	c.regMsg(OpenStateChangeNotify, func() any { return new(proto.OpenStateChangeNotify) })
	c.regMsg(PlayerCookReq, func() any { return new(proto.PlayerCookReq) })
	c.regMsg(PlayerCookRsp, func() any { return new(proto.PlayerCookRsp) })
	c.regMsg(PlayerRandomCookReq, func() any { return new(proto.PlayerRandomCookReq) })
	c.regMsg(PlayerRandomCookRsp, func() any { return new(proto.PlayerRandomCookRsp) })
	c.regMsg(CookDataNotify, func() any { return new(proto.CookDataNotify) })
	c.regMsg(CookRecipeDataNotify, func() any { return new(proto.CookRecipeDataNotify) })
	c.regMsg(CookGradeDataNotify, func() any { return new(proto.CookGradeDataNotify) })
	c.regMsg(PlayerCompoundMaterialReq, func() any { return new(proto.PlayerCompoundMaterialReq) })
	c.regMsg(PlayerCompoundMaterialRsp, func() any { return new(proto.PlayerCompoundMaterialRsp) })
	c.regMsg(TakeCompoundOutputReq, func() any { return new(proto.TakeCompoundOutputReq) })
	c.regMsg(TakeCompoundOutputRsp, func() any { return new(proto.TakeCompoundOutputRsp) })
	c.regMsg(CompoundDataNotify, func() any { return new(proto.CompoundDataNotify) })
	c.regMsg(GetCompoundDataReq, func() any { return new(proto.GetCompoundDataReq) })
	c.regMsg(GetCompoundDataRsp, func() any { return new(proto.GetCompoundDataRsp) })
	c.regMsg(PlayerTimeNotify, func() any { return new(proto.PlayerTimeNotify) })
	c.regMsg(PlayerSetPauseReq, func() any { return new(proto.PlayerSetPauseReq) })
	c.regMsg(PlayerSetPauseRsp, func() any { return new(proto.PlayerSetPauseRsp) })
	c.regMsg(PlayerSetLanguageReq, func() any { return new(proto.PlayerSetLanguageReq) })
	c.regMsg(PlayerSetLanguageRsp, func() any { return new(proto.PlayerSetLanguageRsp) })
	c.regMsg(DataResVersionNotify, func() any { return new(proto.DataResVersionNotify) })
	c.regMsg(DailyTaskDataNotify, func() any { return new(proto.DailyTaskDataNotify) })
	c.regMsg(DailyTaskProgressNotify, func() any { return new(proto.DailyTaskProgressNotify) })
	c.regMsg(DailyTaskScoreRewardNotify, func() any { return new(proto.DailyTaskScoreRewardNotify) })
	c.regMsg(WorldOwnerDailyTaskNotify, func() any { return new(proto.WorldOwnerDailyTaskNotify) })
	c.regMsg(AddRandTaskInfoNotify, func() any { return new(proto.AddRandTaskInfoNotify) })
	c.regMsg(RemoveRandTaskInfoNotify, func() any { return new(proto.RemoveRandTaskInfoNotify) })
	c.regMsg(TakePlayerLevelRewardReq, func() any { return new(proto.TakePlayerLevelRewardReq) })
	c.regMsg(TakePlayerLevelRewardRsp, func() any { return new(proto.TakePlayerLevelRewardRsp) })
	c.regMsg(PlayerLevelRewardUpdateNotify, func() any { return new(proto.PlayerLevelRewardUpdateNotify) })
	c.regMsg(GivingRecordNotify, func() any { return new(proto.GivingRecordNotify) })
	c.regMsg(GivingRecordChangeNotify, func() any { return new(proto.GivingRecordChangeNotify) })
	c.regMsg(ItemGivingReq, func() any { return new(proto.ItemGivingReq) })
	c.regMsg(ItemGivingRsp, func() any { return new(proto.ItemGivingRsp) })
	c.regMsg(PlayerCookArgsReq, func() any { return new(proto.PlayerCookArgsReq) })
	c.regMsg(PlayerCookArgsRsp, func() any { return new(proto.PlayerCookArgsRsp) })
	c.regMsg(PlayerLuaShellNotify, func() any { return new(proto.PlayerLuaShellNotify) })
	c.regMsg(ServerDisconnectClientNotify, func() any { return new(proto.ServerDisconnectClientNotify) })
	c.regMsg(AntiAddictNotify, func() any { return new(proto.AntiAddictNotify) })
	c.regMsg(PlayerForceExitReq, func() any { return new(proto.PlayerForceExitReq) })
	c.regMsg(PlayerForceExitRsp, func() any { return new(proto.PlayerForceExitRsp) })
	c.regMsg(PlayerInjectFixNotify, func() any { return new(proto.PlayerInjectFixNotify) })
	c.regMsg(TaskVarNotify, func() any { return new(proto.TaskVarNotify) })
	c.regMsg(ClientLockGameTimeNotify, func() any { return new(proto.ClientLockGameTimeNotify) })
	c.regMsg(EntityPropNotify, func() any { return new(proto.EntityPropNotify) })
	c.regMsg(LifeStateChangeNotify, func() any { return new(proto.LifeStateChangeNotify) })
	c.regMsg(EntityFightPropNotify, func() any { return new(proto.EntityFightPropNotify) })
	c.regMsg(EntityFightPropUpdateNotify, func() any { return new(proto.EntityFightPropUpdateNotify) })
	c.regMsg(AvatarFightPropNotify, func() any { return new(proto.AvatarFightPropNotify) })
	c.regMsg(AvatarFightPropUpdateNotify, func() any { return new(proto.AvatarFightPropUpdateNotify) })
	c.regMsg(EntityFightPropChangeReasonNotify, func() any { return new(proto.EntityFightPropChangeReasonNotify) })
	c.regMsg(AvatarLifeStateChangeNotify, func() any { return new(proto.AvatarLifeStateChangeNotify) })
	c.regMsg(AvatarPropChangeReasonNotify, func() any { return new(proto.AvatarPropChangeReasonNotify) })
	c.regMsg(PlayerPropChangeReasonNotify, func() any { return new(proto.PlayerPropChangeReasonNotify) })
	c.regMsg(AvatarPropNotify, func() any { return new(proto.AvatarPropNotify) })
	c.regMsg(MarkNewNotify, func() any { return new(proto.MarkNewNotify) })
	c.regMsg(QuestListNotify, func() any { return new(proto.QuestListNotify) })
	c.regMsg(QuestListUpdateNotify, func() any { return new(proto.QuestListUpdateNotify) })
	c.regMsg(QuestDelNotify, func() any { return new(proto.QuestDelNotify) })
	c.regMsg(FinishedParentQuestNotify, func() any { return new(proto.FinishedParentQuestNotify) })
	c.regMsg(FinishedParentQuestUpdateNotify, func() any { return new(proto.FinishedParentQuestUpdateNotify) })
	c.regMsg(AddQuestContentProgressReq, func() any { return new(proto.AddQuestContentProgressReq) })
	c.regMsg(AddQuestContentProgressRsp, func() any { return new(proto.AddQuestContentProgressRsp) })
	c.regMsg(GetQuestTalkHistoryReq, func() any { return new(proto.GetQuestTalkHistoryReq) })
	c.regMsg(GetQuestTalkHistoryRsp, func() any { return new(proto.GetQuestTalkHistoryRsp) })
	c.regMsg(QuestCreateEntityReq, func() any { return new(proto.QuestCreateEntityReq) })
	c.regMsg(QuestCreateEntityRsp, func() any { return new(proto.QuestCreateEntityRsp) })
	c.regMsg(QuestDestroyEntityReq, func() any { return new(proto.QuestDestroyEntityReq) })
	c.regMsg(QuestDestroyEntityRsp, func() any { return new(proto.QuestDestroyEntityRsp) })
	c.regMsg(ChapterStateNotify, func() any { return new(proto.ChapterStateNotify) })
	c.regMsg(QuestProgressUpdateNotify, func() any { return new(proto.QuestProgressUpdateNotify) })
	c.regMsg(QuestUpdateQuestVarReq, func() any { return new(proto.QuestUpdateQuestVarReq) })
	c.regMsg(QuestUpdateQuestVarRsp, func() any { return new(proto.QuestUpdateQuestVarRsp) })
	c.regMsg(QuestUpdateQuestVarNotify, func() any { return new(proto.QuestUpdateQuestVarNotify) })
	c.regMsg(QuestDestroyNpcReq, func() any { return new(proto.QuestDestroyNpcReq) })
	c.regMsg(QuestDestroyNpcRsp, func() any { return new(proto.QuestDestroyNpcRsp) })
	c.regMsg(BargainStartNotify, func() any { return new(proto.BargainStartNotify) })
	c.regMsg(BargainOfferPriceReq, func() any { return new(proto.BargainOfferPriceReq) })
	c.regMsg(BargainOfferPriceRsp, func() any { return new(proto.BargainOfferPriceRsp) })
	c.regMsg(BargainTerminateNotify, func() any { return new(proto.BargainTerminateNotify) })
	c.regMsg(GetBargainDataReq, func() any { return new(proto.GetBargainDataReq) })
	c.regMsg(GetBargainDataRsp, func() any { return new(proto.GetBargainDataRsp) })
	c.regMsg(GetAllActivatedBargainDataReq, func() any { return new(proto.GetAllActivatedBargainDataReq) })
	c.regMsg(GetAllActivatedBargainDataRsp, func() any { return new(proto.GetAllActivatedBargainDataRsp) })
	c.regMsg(ServerCondMeetQuestListUpdateNotify, func() any { return new(proto.ServerCondMeetQuestListUpdateNotify) })
	c.regMsg(QuestGlobalVarNotify, func() any { return new(proto.QuestGlobalVarNotify) })
	c.regMsg(QuestTransmitReq, func() any { return new(proto.QuestTransmitReq) })
	c.regMsg(QuestTransmitRsp, func() any { return new(proto.QuestTransmitRsp) })
	c.regMsg(PersonalLineAllDataReq, func() any { return new(proto.PersonalLineAllDataReq) })
	c.regMsg(PersonalLineAllDataRsp, func() any { return new(proto.PersonalLineAllDataRsp) })
	c.regMsg(RedeemLegendaryKeyReq, func() any { return new(proto.RedeemLegendaryKeyReq) })
	c.regMsg(RedeemLegendaryKeyRsp, func() any { return new(proto.RedeemLegendaryKeyRsp) })
	c.regMsg(UnlockPersonalLineReq, func() any { return new(proto.UnlockPersonalLineReq) })
	c.regMsg(UnlockPersonalLineRsp, func() any { return new(proto.UnlockPersonalLineRsp) })
	c.regMsg(RechargeReq, func() any { return new(proto.RechargeReq) })
	c.regMsg(RechargeRsp, func() any { return new(proto.RechargeRsp) })
	c.regMsg(OrderFinishNotify, func() any { return new(proto.OrderFinishNotify) })
	c.regMsg(CardProductRewardNotify, func() any { return new(proto.CardProductRewardNotify) })
	c.regMsg(PlayerRechargeDataNotify, func() any { return new(proto.PlayerRechargeDataNotify) })
	c.regMsg(OrderDisplayNotify, func() any { return new(proto.OrderDisplayNotify) })
	c.regMsg(PlayerEnterSceneNotify, func() any { return new(proto.PlayerEnterSceneNotify) })
	c.regMsg(LeaveSceneReq, func() any { return new(proto.LeaveSceneReq) })
	c.regMsg(LeaveSceneRsp, func() any { return new(proto.LeaveSceneRsp) })
	c.regMsg(SceneInitFinishReq, func() any { return new(proto.SceneInitFinishReq) })
	c.regMsg(SceneInitFinishRsp, func() any { return new(proto.SceneInitFinishRsp) })
	c.regMsg(SceneEntityAppearNotify, func() any { return new(proto.SceneEntityAppearNotify) })
	c.regMsg(SceneEntityDisappearNotify, func() any { return new(proto.SceneEntityDisappearNotify) })
	c.regMsg(SceneEntityMoveReq, func() any { return new(proto.SceneEntityMoveReq) })
	c.regMsg(SceneEntityMoveRsp, func() any { return new(proto.SceneEntityMoveRsp) })
	c.regMsg(SceneAvatarStaminaStepReq, func() any { return new(proto.SceneAvatarStaminaStepReq) })
	c.regMsg(SceneAvatarStaminaStepRsp, func() any { return new(proto.SceneAvatarStaminaStepRsp) })
	c.regMsg(SceneEntityMoveNotify, func() any { return new(proto.SceneEntityMoveNotify) })
	c.regMsg(ScenePlayerLocationNotify, func() any { return new(proto.ScenePlayerLocationNotify) })
	c.regMsg(GetScenePointReq, func() any { return new(proto.GetScenePointReq) })
	c.regMsg(GetScenePointRsp, func() any { return new(proto.GetScenePointRsp) })
	c.regMsg(EnterTransPointRegionNotify, func() any { return new(proto.EnterTransPointRegionNotify) })
	c.regMsg(ExitTransPointRegionNotify, func() any { return new(proto.ExitTransPointRegionNotify) })
	c.regMsg(ScenePointUnlockNotify, func() any { return new(proto.ScenePointUnlockNotify) })
	c.regMsg(SceneTransToPointReq, func() any { return new(proto.SceneTransToPointReq) })
	c.regMsg(SceneTransToPointRsp, func() any { return new(proto.SceneTransToPointRsp) })
	c.regMsg(EntityJumpNotify, func() any { return new(proto.EntityJumpNotify) })
	c.regMsg(GetSceneAreaReq, func() any { return new(proto.GetSceneAreaReq) })
	c.regMsg(GetSceneAreaRsp, func() any { return new(proto.GetSceneAreaRsp) })
	c.regMsg(SceneAreaUnlockNotify, func() any { return new(proto.SceneAreaUnlockNotify) })
	c.regMsg(SceneEntityDrownReq, func() any { return new(proto.SceneEntityDrownReq) })
	c.regMsg(SceneEntityDrownRsp, func() any { return new(proto.SceneEntityDrownRsp) })
	c.regMsg(SceneCreateEntityReq, func() any { return new(proto.SceneCreateEntityReq) })
	c.regMsg(SceneCreateEntityRsp, func() any { return new(proto.SceneCreateEntityRsp) })
	c.regMsg(SceneDestroyEntityReq, func() any { return new(proto.SceneDestroyEntityReq) })
	c.regMsg(SceneDestroyEntityRsp, func() any { return new(proto.SceneDestroyEntityRsp) })
	c.regMsg(SceneForceUnlockNotify, func() any { return new(proto.SceneForceUnlockNotify) })
	c.regMsg(SceneForceLockNotify, func() any { return new(proto.SceneForceLockNotify) })
	c.regMsg(EnterWorldAreaReq, func() any { return new(proto.EnterWorldAreaReq) })
	c.regMsg(EnterWorldAreaRsp, func() any { return new(proto.EnterWorldAreaRsp) })
	c.regMsg(EntityForceSyncReq, func() any { return new(proto.EntityForceSyncReq) })
	c.regMsg(EntityForceSyncRsp, func() any { return new(proto.EntityForceSyncRsp) })
	c.regMsg(SceneAreaExploreNotify, func() any { return new(proto.SceneAreaExploreNotify) })
	c.regMsg(SceneGetAreaExplorePercentReq, func() any { return new(proto.SceneGetAreaExplorePercentReq) })
	c.regMsg(SceneGetAreaExplorePercentRsp, func() any { return new(proto.SceneGetAreaExplorePercentRsp) })
	c.regMsg(ClientTransmitReq, func() any { return new(proto.ClientTransmitReq) })
	c.regMsg(ClientTransmitRsp, func() any { return new(proto.ClientTransmitRsp) })
	c.regMsg(EnterSceneWeatherAreaNotify, func() any { return new(proto.EnterSceneWeatherAreaNotify) })
	c.regMsg(ExitSceneWeatherAreaNotify, func() any { return new(proto.ExitSceneWeatherAreaNotify) })
	c.regMsg(SceneAreaWeatherNotify, func() any { return new(proto.SceneAreaWeatherNotify) })
	c.regMsg(ScenePlayerInfoNotify, func() any { return new(proto.ScenePlayerInfoNotify) })
	c.regMsg(WorldPlayerLocationNotify, func() any { return new(proto.WorldPlayerLocationNotify) })
	c.regMsg(BeginCameraSceneLookNotify, func() any { return new(proto.BeginCameraSceneLookNotify) })
	c.regMsg(EndCameraSceneLookNotify, func() any { return new(proto.EndCameraSceneLookNotify) })
	c.regMsg(MarkEntityInMinMapNotify, func() any { return new(proto.MarkEntityInMinMapNotify) })
	c.regMsg(UnmarkEntityInMinMapNotify, func() any { return new(proto.UnmarkEntityInMinMapNotify) })
	c.regMsg(DropSubfieldReq, func() any { return new(proto.DropSubfieldReq) })
	c.regMsg(DropSubfieldRsp, func() any { return new(proto.DropSubfieldRsp) })
	c.regMsg(ExecuteGroupTriggerReq, func() any { return new(proto.ExecuteGroupTriggerReq) })
	c.regMsg(ExecuteGroupTriggerRsp, func() any { return new(proto.ExecuteGroupTriggerRsp) })
	c.regMsg(LevelupCityReq, func() any { return new(proto.LevelupCityReq) })
	c.regMsg(LevelupCityRsp, func() any { return new(proto.LevelupCityRsp) })
	c.regMsg(SceneRouteChangeNotify, func() any { return new(proto.SceneRouteChangeNotify) })
	c.regMsg(PlatformStartRouteNotify, func() any { return new(proto.PlatformStartRouteNotify) })
	c.regMsg(PlatformStopRouteNotify, func() any { return new(proto.PlatformStopRouteNotify) })
	c.regMsg(PlatformChangeRouteNotify, func() any { return new(proto.PlatformChangeRouteNotify) })
	c.regMsg(ScenePlayerSoundNotify, func() any { return new(proto.ScenePlayerSoundNotify) })
	c.regMsg(PersonalSceneJumpReq, func() any { return new(proto.PersonalSceneJumpReq) })
	c.regMsg(PersonalSceneJumpRsp, func() any { return new(proto.PersonalSceneJumpRsp) })
	c.regMsg(SealBattleBeginNotify, func() any { return new(proto.SealBattleBeginNotify) })
	c.regMsg(SealBattleEndNotify, func() any { return new(proto.SealBattleEndNotify) })
	c.regMsg(SealBattleProgressNotify, func() any { return new(proto.SealBattleProgressNotify) })
	c.regMsg(ClientPauseNotify, func() any { return new(proto.ClientPauseNotify) })
	c.regMsg(PlayerEnterSceneInfoNotify, func() any { return new(proto.PlayerEnterSceneInfoNotify) })
	c.regMsg(JoinPlayerSceneReq, func() any { return new(proto.JoinPlayerSceneReq) })
	c.regMsg(JoinPlayerSceneRsp, func() any { return new(proto.JoinPlayerSceneRsp) })
	c.regMsg(SceneKickPlayerReq, func() any { return new(proto.SceneKickPlayerReq) })
	c.regMsg(SceneKickPlayerRsp, func() any { return new(proto.SceneKickPlayerRsp) })
	c.regMsg(SceneKickPlayerNotify, func() any { return new(proto.SceneKickPlayerNotify) })
	c.regMsg(HitClientTrivialNotify, func() any { return new(proto.HitClientTrivialNotify) })
	c.regMsg(BackMyWorldReq, func() any { return new(proto.BackMyWorldReq) })
	c.regMsg(BackMyWorldRsp, func() any { return new(proto.BackMyWorldRsp) })
	c.regMsg(SeeMonsterReq, func() any { return new(proto.SeeMonsterReq) })
	c.regMsg(SeeMonsterRsp, func() any { return new(proto.SeeMonsterRsp) })
	c.regMsg(AddSeenMonsterNotify, func() any { return new(proto.AddSeenMonsterNotify) })
	c.regMsg(AllSeenMonsterNotify, func() any { return new(proto.AllSeenMonsterNotify) })
	c.regMsg(SceneTimeNotify, func() any { return new(proto.SceneTimeNotify) })
	c.regMsg(EnterSceneReadyReq, func() any { return new(proto.EnterSceneReadyReq) })
	c.regMsg(EnterSceneReadyRsp, func() any { return new(proto.EnterSceneReadyRsp) })
	c.regMsg(EnterScenePeerNotify, func() any { return new(proto.EnterScenePeerNotify) })
	c.regMsg(EnterSceneDoneReq, func() any { return new(proto.EnterSceneDoneReq) })
	c.regMsg(EnterSceneDoneRsp, func() any { return new(proto.EnterSceneDoneRsp) })
	c.regMsg(WorldPlayerDieNotify, func() any { return new(proto.WorldPlayerDieNotify) })
	c.regMsg(WorldPlayerReviveReq, func() any { return new(proto.WorldPlayerReviveReq) })
	c.regMsg(WorldPlayerReviveRsp, func() any { return new(proto.WorldPlayerReviveRsp) })
	c.regMsg(JoinPlayerFailNotify, func() any { return new(proto.JoinPlayerFailNotify) })
	c.regMsg(SetSceneWeatherAreaReq, func() any { return new(proto.SetSceneWeatherAreaReq) })
	c.regMsg(SetSceneWeatherAreaRsp, func() any { return new(proto.SetSceneWeatherAreaRsp) })
	c.regMsg(ExecuteGadgetLuaReq, func() any { return new(proto.ExecuteGadgetLuaReq) })
	c.regMsg(ExecuteGadgetLuaRsp, func() any { return new(proto.ExecuteGadgetLuaRsp) })
	c.regMsg(CutSceneBeginNotify, func() any { return new(proto.CutSceneBeginNotify) })
	c.regMsg(CutSceneFinishNotify, func() any { return new(proto.CutSceneFinishNotify) })
	c.regMsg(CutSceneEndNotify, func() any { return new(proto.CutSceneEndNotify) })
	c.regMsg(ClientScriptEventNotify, func() any { return new(proto.ClientScriptEventNotify) })
	c.regMsg(SceneEntitiesMovesReq, func() any { return new(proto.SceneEntitiesMovesReq) })
	c.regMsg(SceneEntitiesMovesRsp, func() any { return new(proto.SceneEntitiesMovesRsp) })
	c.regMsg(SceneEntitiesMoveCombineNotify, func() any { return new(proto.SceneEntitiesMoveCombineNotify) })
	c.regMsg(UnlockTransPointReq, func() any { return new(proto.UnlockTransPointReq) })
	c.regMsg(UnlockTransPointRsp, func() any { return new(proto.UnlockTransPointRsp) })
	c.regMsg(PlatformRouteStateNotify, func() any { return new(proto.PlatformRouteStateNotify) })
	c.regMsg(SceneWeatherForcastReq, func() any { return new(proto.SceneWeatherForcastReq) })
	c.regMsg(SceneWeatherForcastRsp, func() any { return new(proto.SceneWeatherForcastRsp) })
	c.regMsg(MarkMapReq, func() any { return new(proto.MarkMapReq) })
	c.regMsg(MarkMapRsp, func() any { return new(proto.MarkMapRsp) })
	c.regMsg(AllMarkPointNotify, func() any { return new(proto.AllMarkPointNotify) })
	c.regMsg(WorldDataNotify, func() any { return new(proto.WorldDataNotify) })
	c.regMsg(EntityMoveRoomNotify, func() any { return new(proto.EntityMoveRoomNotify) })
	c.regMsg(WorldPlayerInfoNotify, func() any { return new(proto.WorldPlayerInfoNotify) })
	c.regMsg(PostEnterSceneReq, func() any { return new(proto.PostEnterSceneReq) })
	c.regMsg(PostEnterSceneRsp, func() any { return new(proto.PostEnterSceneRsp) })
	c.regMsg(PlayerChatReq, func() any { return new(proto.PlayerChatReq) })
	c.regMsg(PlayerChatRsp, func() any { return new(proto.PlayerChatRsp) })
	c.regMsg(PlayerChatNotify, func() any { return new(proto.PlayerChatNotify) })
	c.regMsg(PlayerChatCDNotify, func() any { return new(proto.PlayerChatCDNotify) })
	c.regMsg(ChatHistoryNotify, func() any { return new(proto.ChatHistoryNotify) })
	c.regMsg(SceneDataNotify, func() any { return new(proto.SceneDataNotify) })
	c.regMsg(DungeonEntryToBeExploreNotify, func() any { return new(proto.DungeonEntryToBeExploreNotify) })
	c.regMsg(GetDungeonEntryExploreConditionReq, func() any { return new(proto.GetDungeonEntryExploreConditionReq) })
	c.regMsg(GetDungeonEntryExploreConditionRsp, func() any { return new(proto.GetDungeonEntryExploreConditionRsp) })
	c.regMsg(UnfreezeGroupLimitNotify, func() any { return new(proto.UnfreezeGroupLimitNotify) })
	c.regMsg(SetEntityClientDataNotify, func() any { return new(proto.SetEntityClientDataNotify) })
	c.regMsg(GroupSuiteNotify, func() any { return new(proto.GroupSuiteNotify) })
	c.regMsg(GroupUnloadNotify, func() any { return new(proto.GroupUnloadNotify) })
	c.regMsg(MonsterAIConfigHashNotify, func() any { return new(proto.MonsterAIConfigHashNotify) })
	c.regMsg(EntityAbilityConfigHashNotify, func() any { return new(proto.EntityAbilityConfigHashNotify) })
	c.regMsg(ShowTemplateReminderNotify, func() any { return new(proto.ShowTemplateReminderNotify) })
	c.regMsg(ShowCommonTipsNotify, func() any { return new(proto.ShowCommonTipsNotify) })
	c.regMsg(CloseCommonTipsNotify, func() any { return new(proto.CloseCommonTipsNotify) })
}
