package con02a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	ivar10           int
	firingLineTrig   ns.ObjectID
	wp13             ns.WaypointID
	wp14             ns.WaypointID
	wp15             ns.WaypointID
	wp16             ns.WaypointID
	gvar17           int
	gvar18           int
	gvar19           int
	gvar20           int
	gvar21           int
	gvar22           int
	gvar23           int
	gvar24           int
	gvar25           int
	gvar26           int
	gvar27           int
	gvar28           int
	gvar29           int
	gvar30           int
	gvar31           int
	gvar32           int
	gvar33           int
	gvar34           int
	gvar35           int
	gvar36           int
	gvar37           int
	gvar38           int
	gvar39           int
	gvar40           int
	gvar41           int
	gvar42           int
	gvar43           int
	gvar44           int
	flag45           bool
	flag46           bool
	flag47           bool
	flag48           bool
	flag49           bool
	gvar50           int
	gvar51           int
	gvar52           int
	gvar53           int
	gvar54           int
	gvar55           int
	gvar56           int
	gvar57           int
	contestBowCost   int
	ivar59           int
	ivar60           int
	obj62            ns.ObjectID
	flag63           bool
	flag65           bool
	obj66            ns.ObjectID
	barkeeper        ns.ObjectID
	geoff            ns.ObjectID
	mayorGuard       ns.ObjectID
	aldwyn           ns.ObjectID
	mayor            ns.ObjectID
	leatherTunic     ns.ObjectID
	cloak            ns.ObjectID
	mayorGate1       ns.ObjectID
	mayorGate2       ns.ObjectID
	charmSpell       ns.ObjectID
	ratFieldGuide    ns.ObjectID
	bigSpiderFg      ns.ObjectID
	elevDoor         ns.ObjectID
	aldwynElev       ns.ObjectID
	mayorBedroomDoor ns.ObjectID
	goldObj          ns.ObjectID
	townGateRight    ns.ObjectID
	townGateLeft     ns.ObjectID
	bridgeGuard      ns.ObjectID
	guard2           ns.ObjectID
	bridgeGuardBoots ns.ObjectID
	bootsReward      ns.ObjectID
	morgan           ns.ObjectID
	cheapBow         ns.ObjectID
	magicShopElev    ns.ObjectID
	joyce            ns.ObjectID
	henrick          ns.ObjectID
	mayorFireplace   ns.ObjectID
	gvar103          ns.ObjectGroupID
	wp104            ns.WaypointID
	wp105            ns.WaypointID
	wp106            ns.WaypointID
	wp107            ns.WaypointID
	wp108            ns.WaypointID
	wp109            ns.WaypointID
	wp110            ns.WaypointID
	wp111            ns.WaypointID
	wp112            ns.WaypointID
	wp113            ns.WaypointID
	gvar114          ns.WaypointGroupID
	gvar115          ns.WaypointGroupID
	gvar116          ns.WaypointGroupID
	flag117          bool
	flag118          bool
	obj119           ns.ObjectID
	obj120           ns.ObjectID
	obj121           ns.ObjectID
	pepper           ns.ObjectID
	obj123           ns.ObjectID
	obj124           ns.ObjectID
	bryan            ns.ObjectID
	clyde            ns.ObjectID
	tommy            ns.ObjectID
	gretchen         ns.ObjectID
	obj129           ns.ObjectID
	obj130           ns.ObjectID
	obj131           ns.ObjectID
	obj132           ns.ObjectID
	necroDoor        ns.ObjectID
	gvar134          ns.ObjectGroupID
	groupNPCs        ns.ObjectGroupID
	wp136            ns.WaypointID
	wp137            ns.WaypointID
	wp138            ns.WaypointID
	wp139            ns.WaypointID
	wp140            ns.WaypointID
	wp141            ns.WaypointID
	wp142            ns.WaypointID
	wp143            ns.WaypointID
	wp144            ns.WaypointID
	ivar145          int
	flag146          bool
	flag147          bool
	flag148          bool
	flag149          bool
	gvar150          int
	gvar151          int
	gvar152          int
	gvar153          int
	gvar154          int
	fvar155          float32
	fvar156          float32
	obj157           ns.ObjectID
	flag158          bool
	flag159          bool
	ivar160          int
	wp161            ns.WaypointID
	wp162            ns.WaypointID
	wp163            ns.WaypointID
	wp164            ns.WaypointID
	mayorSpider      ns.ObjectID
	gvar166          ns.ObjectGroupID
	gvar167          ns.ObjectGroupID
	gvar168          ns.WallGroupID
	basementDoor     ns.ObjectID
	gvar170          int
	mayorElev        ns.ObjectID
	gvar172          int
	quiver           ns.ObjectID
	townGate1        ns.ObjectID
	townGate2        ns.ObjectID
	ixGuard1         ns.ObjectID
	ixGuard2         ns.ObjectID
	caveDoor         ns.ObjectID
	waspNest         ns.ObjectID
	wasps            [2]ns.ObjectID
	gvar181          ns.ObjectGroupID
	gvar182          ns.ObjectGroupID
	gvar183          ns.ObjectGroupID
	gvar184          ns.ObjectGroupID
	wp185            ns.WaypointID
	wp186            ns.WaypointID
	wp187            ns.WaypointID
	wp188            ns.WaypointID
	wp189            ns.WaypointID
	wp190            ns.WaypointID
	wp191            ns.WaypointID
	wp192            ns.WaypointID
	wp193            ns.WaypointID
	wp194            ns.WaypointID
	wp195            ns.WaypointID
	wp196            ns.WaypointID
	wp197            ns.WaypointID
	wp198            ns.WaypointID
	wp199            ns.WaypointID
	waspWPs          [2]ns.WaypointID
	warden           ns.ObjectID
	gvar202          int
	gvar203          int
	gvar204          int
	gvar205          int
	gvar206          int
	gvar207          int
	gvar208          int
	gvar209          int
	gvar210          int
	gvar211          int
	gvar212          int
	gvar213          int
	obj214           ns.ObjectID
	obj215           ns.ObjectID
	obj216           ns.ObjectID
	obj217           ns.ObjectID
	gvar218          ns.ObjectGroupID
	wp219            ns.WaypointID
	flag220          bool
	ivar221          int
)

func init() {
	flag4 = true
	contestLastWP = 1
	ivar7 = 0
	ivar8 = 0
	ivar10 = 2
	gvar17 = 0
	gvar18 = 1
	gvar19 = 2
	gvar20 = 3
	gvar21 = 4
	gvar22 = 5
	gvar23 = 6
	gvar24 = 7
	gvar25 = 8
	gvar26 = 9
	gvar27 = 10
	gvar28 = 11
	gvar29 = 12
	gvar30 = 13
	gvar31 = 14
	gvar32 = 15
	gvar33 = 16
	gvar34 = 17
	gvar35 = 18
	gvar36 = 19
	gvar37 = 20
	gvar38 = 21
	gvar39 = 22
	gvar40 = 23
	gvar41 = 24
	gvar42 = 25
	gvar43 = 26
	gvar44 = 27
	flag48 = true
	flag49 = true
	gvar50 = gvar25
	gvar51 = gvar22
	gvar52 = gvar17
	gvar53 = gvar28
	gvar54 = gvar32
	gvar55 = gvar36
	gvar56 = gvar40
	gvar57 = gvar42
	contestBowCost = 20
	ivar59 = 30
	ivar60 = 100
	ivar145 = 0
	gvar150 = 0
	gvar151 = 1
	gvar152 = gvar150
	gvar153 = 0
	gvar154 = 1
	fvar155 = 1650
	fvar156 = 3055
	flag159 = true
	ivar160 = 30
	gvar202 = 0
	gvar203 = 1
	gvar204 = 2
	gvar205 = 3
	gvar206 = 4
	gvar207 = 5
	gvar208 = gvar202
	gvar209 = gvar202
	gvar210 = gvar202
	gvar211 = gvar202
	gvar212 = gvar202
	gvar213 = gvar202
	ivar221 = 0
}

func MapInitialize() {
	aldwyn = ns.Object("Aldwyn")
	geoff = ns.Object("Geoff")
	mayor = ns.Object("Mayor_Theogrin")
	contestGuard = ns.Object("Contest_Guard")
	mayorGuard = ns.Object("Mayor's_Guard")
	barkeeper = ns.Object("Barkeeper")
	contestOfficial = ns.Object("Contest_Official")
	bridgeGuard = ns.Object("Bridge_Guard")
	guard2 = ns.Object("Guard2")
	ixGuard1 = ns.Object("IxGuard1")
	ixGuard2 = ns.Object("IxGuard2")
	mayorSpider = ns.Object("MayorsSpider1")
	mayorGate1 = ns.Object("MayorGate1")
	mayorGate2 = ns.Object("MayorGate2")
	contestGate1 = ns.Object("ContestGate1")
	contestGate2 = ns.Object("ContestGate2")
	contestGate3 = ns.Object("ContestGate3")
	basementDoor = ns.Object("BasementDoor")
	aldwynElev = ns.Object("AldwinElevator")
	mayorElev = ns.Object("MayorElevator")
	ratFieldGuide = ns.Object("RatFieldGuide")
	bigSpiderFg = ns.Object("BigSpiderFG")
	charmSpell = ns.Object("CharmSpell")
	elevDoor = ns.Object("ElevatorDoor")
	cheapBow = ns.Object("CheapBow")
	quiver = ns.Object("Quiver")
	firingLineTrig = ns.Object("FiringLineTrigger")
	contestPrize = ns.Object("ContestPrize")
	contestQuiver1 = ns.Object("ContestQuiver1")
	contestQuiver2 = ns.Object("ContestQuiver2")
	leatherTunic = ns.Object("LeatherTunic")
	mayorBedroomDoor = ns.Object("MayorsBedroomDoor")
	townGateRight = ns.Object("RightTownGate")
	townGateLeft = ns.Object("LeftTownGate")
	townGate1 = ns.Object("TownGate01")
	townGate2 = ns.Object("TownGate02")
	goldObj = ns.Object("Gold")
	bryan = ns.Object("Bryan")
	clyde = ns.Object("Clyde")
	tommy = ns.Object("Tommy")
	cloak = ns.Object("Cloak")
	bridgeGuardBoots = ns.Object("BridgeGuardsBoots")
	bootsReward = ns.Object("BootsReward")
	morgan = ns.Object("Morgan")
	magicShopElev = ns.Object("MagicShopElevator")
	joyce = ns.Object("Joyce")
	gretchen = ns.Object("Gretchen")
	mayorFireplace = ns.Object("MayorsFireplace")
	caveDoor = ns.Object("CaveDoor")
	heckler = ns.Object("Heckler")
	waspNest = ns.Object("WaspNest")
	wasps[0] = ns.Object("Wasp1")
	wasps[1] = ns.Object("Wasp2")
	groupNPCs = ns.ObjectGroup("NPCs")
	gvar181 = ns.ObjectGroup("SpiderDropTriggers1")
	gvar182 = ns.ObjectGroup("SpiderDropTriggers2")
	gvar103 = ns.ObjectGroup("GateTriggers")
	gvar183 = ns.ObjectGroup("WaspTriggers")
	gvar184 = ns.ObjectGroup("Urchins")
	gvar167 = ns.ObjectGroup("Spiders")
	gvar166 = ns.ObjectGroup("Secret01Triggers")
	gvar168 = ns.WallGroup("MagicShopElevatorWalls")
	wp193 = ns.Waypoint("WellWP")
	wp104 = ns.Waypoint("Study")
	wp185 = ns.Waypoint("MayorsGates")
	wp13 = ns.Waypoint("TargetWP1")
	wp14 = ns.Waypoint("TargetWP2")
	wp15 = ns.Waypoint("TargetWP3")
	wp186 = ns.Waypoint("OfficialWP")
	wp187 = ns.Waypoint("CreationWP")
	wp16 = ns.Waypoint("ContestAudioOrigin")
	wp188 = ns.Waypoint("UrchinWP")
	wp105 = ns.Waypoint("ImpWP1")
	wp106 = ns.Waypoint("ImpWP2")
	wp107 = ns.Waypoint("ImpWP3")
	wp108 = ns.Waypoint("ImpWP4")
	wp109 = ns.Waypoint("AldwinExitWP")
	wp189 = ns.Waypoint("SpiderDropPoint1")
	wp190 = ns.Waypoint("SpiderDropPoint2")
	wp110 = ns.Waypoint("MayorRewardPath")
	wp111 = ns.Waypoint("TownGatePath")
	wp112 = ns.Waypoint("MayorHome")
	wp191 = ns.Waypoint("AudioOrigin")
	wp162 = ns.Waypoint("StudySpiderWP")
	wp161 = ns.Waypoint("BigSpiderWP")
	wp163 = ns.Waypoint("SpiderGuardPoint")
	wp192 = ns.Waypoint("WaspCreationWP")
	wp164 = ns.Waypoint("SecretAudioWP")
	wp194 = ns.Waypoint("MayorsGuardWP")
	wp199 = ns.Waypoint("BridgeGuardWP")
	wp198 = ns.Waypoint("ContestGuardWP")
	wp195 = ns.Waypoint("GatekeeperWP")
	wp196 = ns.Waypoint("Gatekeeper2WP")
	wp197 = ns.Waypoint("Gatekeeper3WP")
	waspWPs[0] = ns.Waypoint("WaspWP1")
	waspWPs[1] = ns.Waypoint("WaspWP2")
	gvar114 = ns.WaypointGroup("ImpEntrancePath")
	gvar115 = ns.WaypointGroup("MayorsPathToGate")
	gvar116 = ns.WaypointGroup("MayorsHomeSpiderPath")
	ns.WayPointGroupOff(gvar115)
	ns.LockDoor(mayorGate1)
	ns.LockDoor(mayorGate2)
	ns.LockDoor(townGate1)
	ns.LockDoor(townGate2)
	ns.LockDoor(contestGate1)
	ns.LockDoor(contestGate2)
	ns.LockDoor(contestGate3)
	ns.LockDoor(basementDoor)
	ns.LockDoor(mayorBedroomDoor)
	ns.LockDoor(townGateRight)
	ns.LockDoor(townGateLeft)
	ns.LockDoor(caveDoor)
	InitializePiece()
	InitEvilMage()
	DialogSetup()
	InitializeDialogs()
	ns.StartupScreen(2)
	ns.FrameTimer(1, OwnObjects)
	ns.FrameTimer(35, DisableElevators)
}

func OfficialSEG4() {
	ns.LookAtObject(contestOfficial, ns.GetHost())
	ns.Chat(contestOfficial, "Con02a:OfficialSEG4")
	ContestEngine()
}
func OfficialSEG3() {
	ns.LookAtObject(contestOfficial, ns.GetHost())
	ns.Chat(contestOfficial, "Con02a:OfficialSEG3")
	ns.FrameTimer(30, OfficialSEG4)
}
func OfficialSEG2() {
	ns.LookAtObject(contestOfficial, ns.GetHost())
	ns.ObjectOff(firingLineTrig)
	ns.Chat(contestOfficial, "Con02a:OfficialSEG2")
	ns.FrameTimer(30, OfficialSEG3)
}
func GatekeeperDialogStart() {
	var v0 int
	ns.LookAtObject(geoff, ns.GetHost())
	v0 = gvar56
	if v0 == gvar40 {
		goto LABEL1
	}
	if v0 == gvar41 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:GatekeeperWaiting")
	goto LABEL3
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:GatekeeperGuarding")
	goto LABEL3
LABEL3:
	return
}
func GatekeeperDialogEnd() {
}
func BarkeeperDialogStart() {
	var v0 int
	ns.LookAtObject(barkeeper, ns.GetHost())
	v0 = gvar55
	if v0 == gvar36 {
		goto LABEL1
	}
	if v0 == gvar37 {
		goto LABEL2
	}
	if v0 == gvar38 {
		goto LABEL3
	}
	if v0 == gvar39 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarkeeperGreet")
	goto LABEL5
LABEL2:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarkeeperQuest")
	goto LABEL5
LABEL3:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarkeeperAfterQuest")
	goto LABEL5
LABEL4:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL5
LABEL5:
	return
}
func BarkeeperDialogEnd() {
	if gvar55 != gvar38 {
		goto LABEL1
	}
	gvar55 = gvar39
LABEL1:
	return
}
func MayorsGuardDialogStart() {
	var v0 int
	v0 = gvar50
	if v0 == gvar25 {
		goto LABEL1
	}
	if v0 == gvar26 {
		goto LABEL2
	}
	if v0 == gvar27 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(mayorGuard, ns.GetHost())
	gvar55 = gvar37
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorsGuard")
	goto LABEL4
LABEL2:
	ns.DestroyEveryChat()
	ns.LookAtObject(mayorGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MeetTheMayor")
	if !ns.IsLocked(mayorGate1) {
		goto LABEL5
	}
	ns.UnlockDoor(mayorGate1)
	ns.UnlockDoor(mayorGate2)
LABEL5:
	goto LABEL4
LABEL3:
	ns.LookAtObject(mayorGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "War03b:MayorsGuardIntro")
LABEL4:
	return
}
func MayorsGuardDialogEnd() {
}
func MayorDialogStart() {
	var v0 int
	ns.DestroyEveryChat()
	v0 = gvar53
	if v0 == gvar28 {
		goto LABEL1
	}
	if v0 == gvar29 {
		goto LABEL2
	}
	if v0 == gvar30 {
		goto LABEL3
	}
	if v0 == gvar31 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	ns.LookAtObject(mayor, ns.GetHost())
	ns.DestroyEveryChat()
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorGreeting")
	goto LABEL5
LABEL2:
	ns.LookAtObject(mayor, ns.GetHost())
	if flag117 {
		goto LABEL6
	}
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorProd")
	goto LABEL7
LABEL6:
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorFree")
LABEL7:
	goto LABEL5
LABEL3:
	ns.Frozen(mayor, true)
	ns.LookAtObject(mayor, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorsReward")
	ns.SetShopkeeperText(barkeeper, "Con02a:BarkeeperAfterQuest")
	goto LABEL5
LABEL4:
	ns.LookAtObject(mayor, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:MayorsGuardIdle")
	goto LABEL5
LABEL5:
	return
}
func MayorDialogEnd() {
	var v0 int
	v0 = gvar53
	if v0 == gvar28 {
		goto LABEL1
	}
	if v0 == gvar29 {
		goto LABEL2
	}
	if v0 == gvar30 {
		goto LABEL3
	}
	if v0 == gvar31 {
		goto LABEL4
	}
	goto LABEL5
LABEL1:
	gvar53 = gvar29
	goto LABEL5
LABEL2:
	if !flag117 {
		goto LABEL6
	}
	ns.WayPointGroupOff(gvar116)
	ns.WayPointGroupOn(gvar115)
	ns.UnlockDoor(mayorBedroomDoor)
	ns.CreatureGuard(mayor, ns.GetWaypointX(wp110), ns.GetWaypointY(wp110), ns.GetWaypointX(ns.Waypoint("MayorRewardLookWP")), ns.GetWaypointY(ns.Waypoint("MayorRewardLookWP")), 150)
	gvar53 = gvar30
	ns.CancelDialog(obj66)
	ns.FrameTimer(15, AwardMayorExperience)
LABEL6:
	goto LABEL5
LABEL3:
	ns.Frozen(mayor, false)
	ns.JournalEdit(ns.GetHost(), "BanishSpiders", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.CancelDialog(geoff)
	gvar55 = gvar38
	SaveMayor()
	HelpMayor()
	ns.CancelDialog(mayor)
	gvar50 = gvar27
	ns.Pickup(ns.GetHost(), goldObj)
	ns.Move(mayor, wp111)
	ns.ObjectGroupOn(gvar103)
	if !flag118 {
		goto LABEL7
	}
	ns.PrintToAll("Gate triggers are enabled.")
LABEL7:
	goto LABEL5
LABEL4:
	goto LABEL5
LABEL5:
	return
}
func AldwinDialogStart() {
	var v0 int
	v0 = gvar52
	if v0 == gvar17 {
		goto LABEL1
	}
	if v0 == gvar18 {
		goto LABEL2
	}
	if v0 == gvar19 {
		goto LABEL3
	}
	if v0 == gvar20 {
		goto LABEL4
	}
	if v0 == gvar21 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	ns.LookAtObject(aldwyn, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinGreeting")
	goto LABEL6
LABEL2:
	ns.LookAtObject(aldwyn, ns.GetHost())
	gvar52 = gvar19
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BecomeConjurer2")
	goto LABEL6
LABEL3:
	ns.LookAtObject(aldwyn, ns.GetHost())
	gvar52 = gvar20
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BecomeConjurer3")
	goto LABEL6
LABEL4:
	ns.LookAtObject(aldwyn, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinProd")
	goto LABEL6
LABEL5:
	ns.LookAtObject(aldwyn, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinImp")
	goto LABEL6
LABEL6:
	return
}
func AldwinDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 int
		v6 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v6 = gvar52
	if v6 == gvar17 {
		goto LABEL1
	}
	if v6 == gvar18 {
		goto LABEL2
	}
	if v6 == gvar19 {
		goto LABEL3
	}
	if v6 == gvar20 {
		goto LABEL4
	}
	if v6 == gvar21 {
		goto LABEL5
	}
	goto LABEL6
LABEL1:
	v1 = ns.GetAnswer(aldwyn)
	v5 = v1
	if v5 == v2 {
		goto LABEL7
	}
	if v5 == v3 {
		goto LABEL8
	}
	if v5 == v4 {
		goto LABEL9
	}
	goto LABEL10
LABEL7:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < ivar59) {
		goto LABEL11
	}
	ns.SetDialog(aldwyn, ns.NORMAL, NullDialogStart, ResetAldwinDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:AldwinPoor")
	goto LABEL12
LABEL11:
	if !flag48 {
		goto LABEL13
	}
	flag48 = false
	ns.JournalEdit(ns.GetHost(), "LocateAldwyn", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
LABEL13:
	ns.ChangeGold(ns.GetHost(), -ivar59)
	ns.SetDialog(aldwyn, ns.NEXT, AldwinDialogStart, AldwinDialogEnd)
	gvar52 = gvar18
	flag45 = true
	gvar50 = gvar26
	ns.Pickup(ns.GetHost(), charmSpell)
	ns.ObjectOn(aldwynElev)
	ns.ObjectOn(magicShopElev)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BecomeConjurer1")
	ns.JournalEntry(ns.GetHost(), "BanishSpiders", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL12:
	goto LABEL10
LABEL8:
	goto LABEL10
LABEL9:
	goto LABEL10
LABEL10:
	goto LABEL6
LABEL2:
	gvar52 = gvar18
	AldwinDialogStart()
	ns.Pickup(ns.GetHost(), bigSpiderFg)
	goto LABEL6
LABEL3:
	ns.SetDialog(aldwyn, ns.NORMAL, AldwinDialogStart, AldwinDialogEnd)
	ns.GiveXp(ns.GetHost(), 200)
	AldwinDialogStart()
	goto LABEL6
LABEL4:
	ns.CreatureGuard(aldwyn, ns.GetWaypointX(wp104), ns.GetWaypointY(wp104), ns.GetObjectX(ns.Object("AldwynBookcase")), ns.GetObjectY(ns.Object("AldwynBookcase")), 0)
	goto LABEL6
LABEL5:
	ns.JournalEntry(ns.GetHost(), "LocateMineForeman", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.SetDialog(geoff, ns.NORMAL, GatekeeperDialogStart, GatekeeperDialogEnd)
	ns.Frozen(ns.GetHost(), false)
	ns.UnlockDoor(townGateRight)
	ns.UnlockDoor(townGateLeft)
	ns.CancelDialog(aldwyn)
	ns.FrameTimer(30, ImpExit)
	goto LABEL6
LABEL6:
	return
}
func CheckForAldwyn() {
	if !ns.IsCaller(aldwyn) {
		goto LABEL1
	}
	ns.CreatureGuard(aldwyn, ns.GetWaypointX(wp104), ns.GetWaypointY(wp104), 5117, 1091, 0)
LABEL1:
	return
}
func BridgeGuardDialogStart() {
	var (
		v0 ns.ObjectID
		v1 int
	)
	v0 = ns.GetLastItem(ns.GetHost())
	for {
		if v0 == 0 {
			goto LABEL1
		}
		if v0 != bridgeGuardBoots {
			goto LABEL2
		}
		ns.Delete(v0)
		flag65 = true
		goto LABEL1
	LABEL2:
		v0 = ns.GetPreviousItem(v0)
	}
LABEL1:
	if !(flag65 == true && flag47 == false) {
		goto LABEL4
	}
	gvar57 = gvar43
LABEL4:
	v1 = gvar57
	if v1 == gvar42 {
		goto LABEL5
	}
	if v1 == gvar43 {
		goto LABEL6
	}
	if v1 == gvar44 {
		goto LABEL7
	}
	goto LABEL8
LABEL5:
	ns.LookAtObject(bridgeGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BridgeGuardGuarding")
	goto LABEL8
LABEL6:
	ns.LookAtObject(bridgeGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BridgeGuardReward")
	ns.Drop(ns.GetHost(), bridgeGuardBoots)
	ns.FrameTimer(1, PickupBoots)
	flag47 = true
	goto LABEL8
LABEL7:
	ns.LookAtObject(bridgeGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL8
LABEL8:
	return
}
func BridgeGuardDialogEnd() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
		v4 int
	)
	v4 = gvar57
	if v4 == gvar42 {
		goto LABEL1
	}
	if v4 == gvar43 {
		goto LABEL2
	}
	if v4 == gvar44 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	if !flag49 {
		goto LABEL5
	}
	flag49 = false
	ns.JournalEntry(ns.GetHost(), "ReturnBridgeGuardsBoots", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
LABEL5:
	goto LABEL4
LABEL2:
	ns.CancelDialog(bridgeGuard)
	v0 = ns.GetObjectX(bridgeGuard)
	v1 = ns.GetObjectY(bridgeGuard)
	v2 = 3116
	v3 = 3394
	ns.MoveObject(bridgeGuard, v2, v3)
	ns.MoveObject(guard2, v0, v1)
	ns.CreatureGuard(bridgeGuard, v2, v3, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(guard2, v0, v1, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.LookAtObject(guard2, ns.GetHost())
	ns.JournalEdit(ns.GetHost(), "ReturnBridgeGuardsBoots", 4)
	ns.PrintToAll("Con02a:ObjectiveComplete")
	ns.FrameTimer(15, AwardBootQuestExperience)
	goto LABEL4
LABEL3:
	goto LABEL4
LABEL4:
	return
}
func ConManDialogStart() {
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ConManSalesPitch")
}
func ConManDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
	)
	v0 = 0
	v1 = 1
	v2 = 0
	v0 = ns.GetAnswer(morgan)
	v2 = ns.GetGold(ns.GetHost())
	if !ns.HasItem(ns.GetHost(), cheapBow) {
		goto LABEL1
	}
	ns.CancelDialog(morgan)
	return
LABEL1:
	if v0 != v1 {
		goto LABEL2
	}
	if !(v2 < ivar60) {
		goto LABEL3
	}
	ns.SetDialog(morgan, ns.NORMAL, NullDialogStart, ResetConManDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ConManNotEnoughGold")
	goto LABEL2
LABEL3:
	ns.Pickup(ns.GetHost(), cheapBow)
	ns.PrintToAll("Con02a:BowAddedToInventory")
	ns.ChangeGold(ns.GetHost(), -ivar60)
	ns.SetDialog(morgan, ns.NORMAL, NullDialogStart, ConManDialogEnd)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ConManSale")
	SlipIntoShadows()
LABEL2:
	return
}
func Maiden1DialogStart() {
	ns.LookAtObject(joyce, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:BarMaiden")
}
func Maiden1DialogEnd() {
}
func SpiderEngine() {
	var v0 int
	if !flag118 {
		goto LABEL1
	}
	ns.PrintToAll("SpiderEngineRunning")
LABEL1:
	v0 = gvar152
	if v0 == gvar150 {
		goto LABEL2
	}
	if v0 == gvar151 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	if !(ns.IsOwnedBy(mayorSpider, ns.GetHost()) || ns.CurrentHealth(mayorSpider) <= 0) {
		goto LABEL5
	}
	obj157 = ns.CreateObject("Spider", wp161)
	ns.SetCallback(obj157, 5, TransientMonsterDie)
	ns.LookAtObject(obj157, ns.GetHost())
	gvar152 = gvar151
	ns.FrameTimerWithArg(1, gvar151, SetToGuard)
	ns.FrameTimer(30, SpiderEngine)
	goto LABEL6
LABEL5:
	ns.FrameTimer(15, SpiderEngine)
LABEL6:
	goto LABEL4
LABEL3:
	if !(ns.CurrentHealth(obj157) > 0 || ns.CurrentHealth(mayorSpider) > 0) {
		goto LABEL7
	}
	ns.FrameTimer(15, SpiderEngine)
	goto LABEL8
LABEL7:
	flag117 = true
	SpidersFinishedPatch()
LABEL8:
	goto LABEL4
LABEL4:
	return
}
func HelpMayor() {
	gvar208 = gvar204
	gvar209 = gvar204
	gvar210 = gvar204
	gvar212 = gvar204
	gvar213 = gvar204
}
func SaveMayor() {
	flag147 = true
}
func ImpEnter() {
	if !flag118 {
		goto LABEL1
	}
	ns.PrintToAll("ImpEnter()")
LABEL1:
	ns.ObjectGroupOff(gvar103)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	obj62 = ns.CreateObject("Imp", wp105)
	ns.SetOwner(ns.GetHost(), obj62)
	ns.Enchant(obj62, ns.ENCHANT_INVULNERABLE, 0)
	ns.FrameTimer(1, ImpEnterSEG2)
}
func AldwinWaiting() {
	if !ns.IsTalking() {
		goto LABEL1
	}
	ns.FrameTimer(1, AldwinWaiting)
	goto LABEL2
LABEL1:
	ns.StartDialog(aldwyn, ns.GetHost())
LABEL2:
	return
}
func InitializeDialogs() {
	ns.StoryPic(geoff, "Townsman4Pic")
	ns.StoryPic(contestGuard, "Townsman2Pic")
	ns.StoryPic(contestOfficial, "MalePic10")
	ns.StoryPic(mayorGuard, "Warrior3Pic")
	ns.StoryPic(mayor, "TheogrinPic")
	ns.StoryPic(aldwyn, "AldwynPic")
	ns.StoryPic(bridgeGuard, "Townsman3Pic")
	ns.StoryPic(morgan, "MorganPic")
	ns.StoryPic(joyce, "MaidenPic")
	ns.SetDialog(geoff, ns.NORMAL, GatekeeperDialogStart, GatekeeperDialogEnd)
	ns.SetDialog(contestGuard, ns.YESNO, ContestGuardDialogStart, ContestGuardDialogEnd)
	ns.SetDialog(mayorGuard, ns.NORMAL, MayorsGuardDialogStart, MayorsGuardDialogEnd)
	ns.SetDialog(mayor, ns.NORMAL, MayorDialogStart, MayorDialogEnd)
	ns.SetDialog(aldwyn, ns.YESNO, AldwinDialogStart, AldwinDialogEnd)
	ns.SetDialog(contestOfficial, ns.NORMAL, ContestOfficialDialogStart, ContestOfficialDialogEnd)
	ns.SetDialog(bridgeGuard, ns.NORMAL, BridgeGuardDialogStart, BridgeGuardDialogEnd)
	ns.SetDialog(morgan, ns.YESNO, ConManDialogStart, ConManDialogEnd)
	ns.SetDialog(joyce, ns.NORMAL, Maiden1DialogStart, Maiden1DialogEnd)
}
func NullDialogStart() {
}
func AwardBootQuestExperience() {
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.FlagDrop, wp113)
	ns.GiveXp(ns.GetHost(), 200)
}
func AwardContestExperience() {
	ns.GiveXp(ns.GetHost(), 500)
}
func AwardMayorExperience() {
	ns.GiveXp(ns.GetHost(), 1500)
}
func PickupBoots() {
	ns.Pickup(bridgeGuard, bridgeGuardBoots)
	ns.Pickup(ns.GetHost(), bootsReward)
}
func EnableFiringLineTrigger() {
	ns.ObjectOn(firingLineTrig)
	ns.SetDialog(contestOfficial, ns.NORMAL, ContestOfficialDialogStart, ContestOfficialDialogEnd)
}
func ImpToAldwin() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetWaypointX(wp106)
	v1 = ns.GetWaypointY(wp106)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Delete(obj62)
	ns.MoveObject(aldwyn, v0, v1)
	gvar52 = gvar21
	ns.LookAtObject(aldwyn, ns.GetHost())
	ns.AudioEvent(ns.ImpRecognize, wp106)
	ns.CreatureGuard(aldwyn, ns.GetObjectX(aldwyn), ns.GetObjectY(aldwyn), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	AldwinWaiting()
}
func ImpEnterSEG2() {
	ns.Move(obj62, wp106)
	ns.FrameTimer(75, ImpToAldwin)
}
func DestroyImp() {
	var (
		v0 float32
		v1 float32
	)
	v0 = ns.GetObjectX(obj62)
	v1 = ns.GetObjectY(obj62)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.Delete(obj62)
}
func ImpExitSEG2() {
	ns.Move(obj62, wp107)
	ns.FrameTimer(75, DestroyImp)
}
func ImpExit() {
	var (
		v0 float32
		v1 float32
		v2 float32
		v3 float32
	)
	v0 = ns.GetObjectX(aldwyn)
	v1 = ns.GetObjectY(aldwyn)
	v2 = ns.GetWaypointX(wp109)
	v3 = ns.GetWaypointY(wp109)
	ns.Effect(ns.SMOKE_BLAST, v0, v1, 0, 0)
	ns.MoveObject(aldwyn, v2, v3)
	ns.ObjectOff(aldwyn)
	obj62 = ns.CreateObject("Imp", wp106)
	ns.AudioEvent(ns.ImpRecognize, wp106)
	ns.SetOwner(ns.GetHost(), obj62)
	ns.Enchant(obj62, ns.ENCHANT_INVULNERABLE, 0)
	ns.WayPointGroupOff(gvar114)
	ns.FrameTimer(1, ImpExitSEG2)
}
func ResetAldwinDialogChoice() {
	ns.SetDialog(aldwyn, ns.YESNO, AldwinDialogStart, AldwinDialogEnd)
}
func SpidersFinishedPatch() {
	gvar53 = gvar29
}
func SlipIntoShadows() {
	if ns.IsVisibleTo(ns.GetHost(), morgan) {
		goto LABEL1
	}
	ns.MoveObject(morgan, ns.GetWaypointX(ns.Waypoint("ConManEscapeWP")), ns.GetWaypointY(ns.Waypoint("ConManEscapeWP")))
	ns.ObjectOff(morgan)
	goto LABEL2
LABEL1:
	ns.FrameTimer(15, SlipIntoShadows)
LABEL2:
	return
}
func ResetConManDialogChoice() {
	ns.SetDialog(morgan, ns.YESNO, ConManDialogStart, ConManDialogEnd)
}
func MayorGoHome() {
	ns.CreatureGuard(mayor, ns.GetWaypointX(wp112), ns.GetWaypointY(wp112), ns.GetObjectX(mayorFireplace), ns.GetObjectY(mayorFireplace), 0)
}
func UnlockTownGates() {
	gvar56 = gvar41
	ns.FrameTimer(30, MayorGoHome)
}
func PlayerDeath() {
	ns.DeathScreen(2)
}
func SwitchMaps() {
	ns.MoveObject(ns.GetHost(), ns.GetWaypointX(ns.Waypoint("ExitWP")), ns.GetWaypointY(ns.Waypoint("ExitWP")))
}
func EndMission() {
	ns.Blind()
	ns.FrameTimer(60, SwitchMaps)
}
func InitEvilMage() {
	obj119 = ns.Object("Necromancer")
	obj123 = ns.Object("Wolf1")
	obj124 = ns.Object("Wolf2")
	obj129 = ns.Object("Julie2")
	obj130 = ns.Object("Tanya2")
	obj131 = ns.Object("Julie")
	obj66 = ns.Object("Tanya")
	obj132 = ns.Object("Lydia")
	obj121 = ns.Object("SummonTrigger")
	gvar134 = ns.ObjectGroup("NecroTriggers")
	wp136 = ns.Waypoint("NecroWP")
	wp137 = ns.Waypoint("SpiderWP")
	wp138 = ns.Waypoint("EscapeWP")
	wp139 = ns.Waypoint("T1WP")
	wp140 = ns.Waypoint("T2WP")
	wp141 = ns.Waypoint("M3WP")
	wp142 = ns.Waypoint("M4WP")
	wp143 = ns.Waypoint("M3RunWP")
	wp144 = ns.Waypoint("M4RunWP")
	ns.ObjectOff(obj121)
	ns.FrameTimer(1, FunctionSetup)
}
func FunctionSetup() {
	ns.SetOwner(ns.GetHost(), obj129)
	ns.SetOwner(ns.GetHost(), obj130)
	ns.SetOwner(ns.GetHost(), obj131)
	ns.SetOwner(ns.GetHost(), obj66)
	ns.SetOwner(ns.GetHost(), obj132)
	ns.SetOwner(ns.GetHost(), obj123)
	ns.SetOwner(ns.GetHost(), obj124)
	ns.StoryPic(obj130, "MaidenPic3")
	ns.StoryPic(obj132, "SnobbyGirlPic")
	ns.SetDialog(obj130, ns.NORMAL, Maiden4TempStart, Maiden4TempEnd)
	ns.SetDialog(obj132, ns.NORMAL, Maiden7TempStart, Maiden7TempEnd)
}
func T1Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145++
	ns.CreatureGuard(bryan, ns.GetObjectX(bryan), ns.GetObjectY(bryan), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func T2Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145++
	ns.CreatureGuard(ns.GetTrigger(), ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func M3Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145++
	ns.CreatureGuard(ns.GetTrigger(), ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func M4Trigger() {
	if flag146 {
		goto LABEL1
	}
	ivar145++
	ns.CreatureGuard(ns.GetTrigger(), ns.GetObjectX(ns.GetTrigger()), ns.GetObjectY(ns.GetTrigger()), ns.GetObjectX(obj119), ns.GetObjectY(obj119), 0)
	if !(ivar145 >= 4) {
		goto LABEL1
	}
	NecromancerSummon()
LABEL1:
	return
}
func StartNecroPiece() {
	ns.MusicPushEvent()
	ns.ObjectGroupOff(gvar134)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CancelDialog(obj130)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), ns.ObjectID(wp136)) // FIXME
	ns.Move(obj119, wp136)
	ns.ObjectOn(bryan)
	ns.ObjectOn(clyde)
}
func NecroInPosition() {
	ns.CreatureIdle(obj119)
	ns.LookAtObject(obj119, ns.GetHost())
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.NecromancerRecognize, wp113)
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.MaleNPC1Talkable, wp113)
	ns.AudioEvent(ns.MaleNPC2Talkable, wp113)
	ns.Move(bryan, wp139)
	ns.Move(clyde, wp140)
	ns.Move(obj129, wp141)
	ns.Move(obj130, wp142)
	ns.FrameTimer(1, NecromancerStart)
}
func NecromancerStart() {
	ns.Music(11, 100)
	ns.SetDialog(obj119, ns.NORMAL, NecroBegin, NecroDone)
	ns.StoryPic(obj119, "NecromancerPic")
	ns.StartDialog(obj119, ns.GetHost())
}
func NecromancerSummon() {
	if !flag149 {
		goto LABEL1
	}
	flag146 = true
	ns.ObjectOn(obj121)
	ns.LookAtObject(obj119, obj121)
	ns.CastSpellObjectLocation(ns.SPELL_SUMMON_SPIDER, obj119, 1374, 3207)
	ns.AudioEvent(ns.NecromancerTaunt, wp113)
	goto LABEL2
LABEL1:
	ns.FrameTimer(2, NecromancerSummon)
LABEL2:
	return
}
func SpiderRun() {
	r1 := ns.IsAttackedBy(ns.GetCaller(), ns.GetHost())
	if !r1 {
		goto LABEL1
	}
	if ns.IsCaller(obj119) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.AudioEvent(ns.NecromancerTaunt, wp113)
	obj120 = ns.GetCaller()
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.Maiden1Retreat, wp113)
	ns.Move(obj129, wp143)
	ns.FrameTimer(1, SpiderActivate)
LABEL1:
	return
}
func SpiderActivate() {
	ns.AggressionLevel(obj120, 0)
	ns.SetOwner(ns.GetHost(), obj120)
	ns.Move(obj120, wp137)
	ns.AudioEvent(ns.Maiden2Retreat, wp113)
	ns.Move(obj130, wp144)
	HenrickSayAttack()
}
func HenrickSayAttack() {
	ns.FrameTimer(15, PepperAttack)
}
func PepperAttack() {
	ns.AudioEvent(ns.SwordsmanRecognize, wp113)
	ns.Attack(clyde, obj119)
	ns.CreatureGuard(obj119, ns.GetObjectX(obj119), ns.GetObjectY(obj119), ns.GetObjectX(clyde), ns.GetObjectY(clyde), 0)
}
func NecroHit() {
	if flag148 {
		goto LABEL1
	}
	if !ns.IsCaller(clyde) {
		goto LABEL1
	}
	flag148 = true
	ns.Chat(obj119, "Con02A:NecroTalk02")
	ns.FrameTimer(60, NecroPoof)
LABEL1:
	return
}
func NecroPoof() {
	ns.DestroyEveryChat()
	ns.CreatureIdle(bryan)
	ns.CreatureIdle(clyde)
	ns.LookAtObject(clyde, obj119)
	ns.LookAtObject(bryan, obj119)
	ns.Move(obj119, wp138)
	ns.Attack(clyde, obj119)
	ns.Attack(bryan, obj119)
	ns.Chat(obj119, "Con02A:NecroTalk03")
	ns.FrameTimer(120, SetpieceOver)
}
func SetpieceOver() {
	ns.DestroyEveryChat()
	ns.WideScreen(false)
	ns.ObjectGroupOn(groupNPCs)
	ns.Frozen(ns.GetHost(), false)
	ns.Delete(obj119)
	ns.MoveObject(bryan, 1678, 2623)
	ns.MoveObject(clyde, 1708, 2655)
	ns.CreatureGuard(bryan, 1678, 2623, 1708, 2655, 0)
	ns.CreatureGuard(clyde, 1708, 2655, 1678, 2623, 0)
	ns.Wander(tommy)
	ns.Wander(gretchen)
	ns.Wander(obj131)
	ns.Wander(obj66)
	ns.StoryPic(obj131, "MaidenPic2")
	ns.StoryPic(obj66, "MaidenPic3")
	ns.SetDialog(bryan, ns.NORMAL, Town1TempStart, Town1TempEnd)
	ns.SetDialog(clyde, ns.NORMAL, Town2TempStart, Town2TempEnd)
	ns.SetDialog(obj131, ns.NORMAL, Maiden5TempStart, Maiden5TempEnd)
	ns.SetDialog(obj66, ns.NORMAL, Maiden6TempStart, Maiden6TempEnd)
	ns.MusicPopEvent()
}
func DestroySpider() {
	if !ns.IsCaller(obj120) {
		goto LABEL1
	}
	ns.Delete(ns.GetCaller())
LABEL1:
	return
}
func M3Swap() {
	if !ns.IsCaller(obj129) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj131, ns.GetObjectX(obj129), ns.GetObjectY(obj129))
	ns.Delete(obj129)
LABEL1:
	return
}
func M4Swap() {
	if !ns.IsCaller(obj130) {
		goto LABEL1
	}
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveObject(obj66, ns.GetObjectX(obj130), ns.GetObjectY(obj130))
	ns.Delete(obj130)
LABEL1:
	return
}
func NecroBegin() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:NecroTalk01")
}
func NecroDone() {
	ns.CancelDialog(obj119)
	ns.CreatureGuard(obj119, ns.GetObjectX(obj119), ns.GetObjectY(obj119), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	flag149 = true
}
func NecroYellStart() {
	ns.Frozen(bryan, true)
	ns.Frozen(clyde, true)
	ns.CreatureIdle(bryan)
	ns.CreatureIdle(clyde)
	ns.LookAtObject(clyde, obj119)
	ns.LookAtObject(bryan, obj119)
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:NecroTalk03")
}
func NecroYellEnd() {
	ns.CancelDialog(obj119)
	ns.Frozen(clyde, false)
	ns.Frozen(bryan, false)
	ns.Move(obj119, wp138)
	ns.Attack(clyde, obj119)
	ns.Attack(bryan, obj119)
	ns.FrameTimer(120, SetpieceOver)
}
func Town1TempStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk03")
}
func Town1TempEnd() {
	ns.Wander(bryan)
	ns.SetDialog(bryan, ns.NORMAL, Town1TalkStart, Town1TalkEnd)
}
func Town2TempStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk02")
}
func Town2TempEnd() {
	ns.Wander(clyde)
	ns.SetDialog(clyde, ns.NORMAL, Town2TalkStart, Town2TalkEnd)
}
func Maiden5TempStart() {
	ns.Frozen(obj131, true)
	ns.CreatureIdle(obj131)
	ns.LookAtObject(obj131, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk03")
}
func Maiden5TempEnd() {
	ns.Frozen(obj131, false)
	ns.Wander(obj131)
	ns.SetDialog(obj131, ns.NORMAL, Maiden5TalkStart, Maiden5TalkEnd)
}
func Maiden4TempStart() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden3Talk01")
}
func Maiden4TempEnd() {
	ns.SetDialog(obj130, ns.NORMAL, Maiden4TempStart, Maiden4TempEnd)
}
func Maiden6TempStart() {
	ns.Frozen(obj66, true)
	ns.CreatureIdle(obj66)
	ns.LookAtObject(obj66, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk01")
}
func Maiden6TempEnd() {
	ns.Frozen(obj66, false)
	ns.Wander(obj66)
	ns.SetDialog(obj66, ns.NORMAL, Maiden6TempStart, Maiden6TempEnd)
}
func Maiden7TempStart() {
	ns.LookAtObject(obj132, ns.GetHost())
	ns.TellStory(ns.HumanMaleEatFood, "Con02a:BarMaiden2")
}
func Maiden7TempEnd() {
	ns.SetDialog(obj132, ns.NORMAL, Maiden7TempStart, Maiden7TempEnd)
}
func Town1TalkStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(bryan, true)
	ns.CreatureIdle(bryan)
	ns.LookAtObject(bryan, ns.GetHost())
	if gvar208 != gvar202 {
		goto LABEL1
	}
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL2
	}
	if v1 == 2 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk02")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk03")
	goto LABEL4
LABEL4:
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman10Talk01")
LABEL5:
	return
}
func Town1TalkEnd() {
	ns.Frozen(bryan, false)
	ns.Wander(bryan)
	ns.SetDialog(bryan, ns.NORMAL, Town1TalkStart, Town1TalkEnd)
}
func Town2TalkStart() {
	var (
		v0 int
		v1 int
	)
	ns.Frozen(clyde, true)
	ns.CreatureIdle(clyde)
	ns.LookAtObject(clyde, ns.GetHost())
	if gvar209 != gvar204 {
		goto LABEL1
	}
	v0 = ns.Random(1, 2)
	v1 = v0
	if v1 == 1 {
		goto LABEL2
	}
	if v1 == 2 {
		goto LABEL3
	}
	goto LABEL4
LABEL2:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk01")
	goto LABEL4
LABEL3:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk03")
	goto LABEL4
LABEL4:
	goto LABEL5
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman13Talk02")
LABEL5:
	return
}
func Town2TalkEnd() {
	ns.Frozen(clyde, false)
	ns.Wander(clyde)
	ns.SetDialog(clyde, ns.NORMAL, Town2TalkStart, Town2TalkEnd)
}
func Maiden5TalkStart() {
	ns.Frozen(obj131, true)
	ns.CreatureIdle(obj131)
	ns.LookAtObject(obj131, ns.GetHost())
	if flag147 {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk03")
	goto LABEL2
LABEL1:
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden6Talk02")
LABEL2:
	return
}
func Maiden5TalkEnd() {
	ns.Frozen(obj131, false)
	ns.Wander(obj131)
	ns.SetDialog(obj131, ns.NORMAL, Maiden5TalkStart, Maiden5TalkEnd)
}
func EnableRandomBump() {
	flag159 = true
}
func RandomBump() {
	if !(ns.IsCaller(ns.GetHost()) && flag159) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:RandomBump")
	flag159 = false
	ns.FrameTimer(ivar160, EnableRandomBump)
LABEL1:
	return
}
func RestAwhile() {
	var v0 int
	v0 = ns.Random(60, 120)
	ns.PauseObject(ns.GetCaller(), v0)
}
func EnableSpiders() {
	ns.ObjectGroupOn(gvar167)
}
func OpenMagicShopElevatorWalls() {
	ns.WallGroupOpen(gvar168)
}
func TransientMonsterDie() {
	ns.DeleteObjectTimer(ns.GetTrigger(), 120)
}
func SetToGuard(a1 int) {
	var (
		v0 float32
		v1 float32
		v2 int
	)
	v2 = a1
	if v2 == gvar150 {
		goto LABEL1
	}
	if v2 == gvar151 {
		goto LABEL2
	}
	goto LABEL3
LABEL1:
	v0 = ns.GetWaypointX(wp163)
	v1 = ns.GetWaypointY(wp163)
	ns.MoveObject(mayorSpider, ns.GetWaypointX(wp161), ns.GetWaypointY(wp161))
	ns.CreatureGuard(mayorSpider, v0, v1, fvar155, fvar156, 100)
	goto LABEL3
LABEL2:
	v0 = ns.GetWaypointX(wp163)
	v1 = ns.GetWaypointY(wp163)
	ns.CreatureGuard(obj157, v0, v1, fvar155, fvar156, 100)
	goto LABEL3
LABEL3:
	return
}
func CallToConjurer() {
	if !(ns.IsCaller(ns.GetHost()) && flag45 == true && flag158 == false) {
		goto LABEL1
	}
	ns.Chat(ns.GetTrigger(), "Con02a:CallToConjurer")
	flag158 = true
LABEL1:
	return
}
func MayorCall() {
	ns.Chat(mayor, "Con02a:MayorCall")
	SpiderEngine()
}
func MayorsGateTriggerFilter() int {
	if !ns.IsCaller(mayor) {
		goto LABEL1
	}
	if !flag118 {
		goto LABEL2
	}
	ns.PrintToAll("Mayor detected at gown gates.")
LABEL2:
	return 1
	goto LABEL3
LABEL1:
	if !flag118 {
		goto LABEL4
	}
	ns.PrintToAll("Trigger hit by object other than Mayor Theogrin")
LABEL4:
	return 0
LABEL3:
	return 0
}
func UnlockMayorBedroomoor() {
	ns.UnlockDoor(mayorBedroomDoor)
}
func LockMayorsBedroomDoor() {
	ns.LockDoor(mayorBedroomDoor)
}
func FoundBoots() {
	ns.ObjectOff(ns.GetTrigger())
	ns.PrintToAll("Con02a:FoundBoots")
	ns.AudioEvent(ns.JournalEntryAdd, ns.Waypoint("BootsAudioOrigin"))
}
func Secret01Found() {
	ns.ObjectGroupOff(gvar166)
	ns.MoveWaypoint(wp164, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp164)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func Secret02Found() {
	ns.ObjectOff(ns.GetTrigger())
	ns.MoveWaypoint(wp164, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.SecretFound, wp164)
	ns.PrintToAll("GeneralPrint:SecretFound")
	ns.GiveXp(ns.GetHost(), 50)
}
func EnemyGoHome() {
	r1 := ns.IsAttackedBy(ns.GetHost(), ns.GetCaller())
	if !r1 {
		goto LABEL1
	}
	ns.GoBackHome(ns.GetCaller())
LABEL1:
	return
}
func DebugOn() {
	flag118 = true
	ns.PrintToAll("Debug mode is enabled.")
}
func OwnObjects() {
	ns.GroupSetOwner(ns.GetHost(), groupNPCs)
	ns.SetOwner(ns.GetHost(), geoff)
	ns.SetOwner(ns.GetHost(), morgan)
	ns.SetOwner(ns.GetHost(), joyce)
	ns.SetOwner(ns.GetHost(), gretchen)
	ns.SetOwner(ns.GetHost(), ixGuard1)
	ns.SetOwner(ns.GetHost(), ixGuard2)
	ns.SetOwner(ns.GetHost(), bryan)
	ns.SetOwner(ns.GetHost(), clyde)
	ns.SetOwner(ns.GetHost(), tommy)
	ns.SetOwner(ns.GetHost(), guard2)
	ns.ObjectGroupOff(groupNPCs)
}
func InitializePiece() {
	warden = ns.Object("Jacob")
	henrick = ns.Object("Henrick")
	pepper = ns.Object("Pepper")
	necroDoor = ns.Object("NecroDoor")
	obj214 = ns.Object("TempDoorR")
	obj215 = ns.Object("TempDoorL")
	obj216 = ns.Object("CemDoorR")
	obj217 = ns.Object("CemDoorL")
	gvar218 = ns.ObjectGroup("WolfTriggers")
	wp113 = ns.Waypoint("PlayerSounds")
	wp219 = ns.Waypoint("WolfCreateWP")
	ns.FrameTimer(1, SetupCreatures)
}
func DialogSetup() {
	ns.StoryPic(bryan, "MalePic4")
	ns.StoryPic(clyde, "MalePic11")
	ns.StoryPic(tommy, "MalePic12")
	ns.StoryPic(gretchen, "MaidenPic4")
	ns.StoryPic(warden, "WardenPic")
	ns.StoryPic(henrick, "HenrickPic")
	ns.SetDialog(tommy, ns.NORMAL, Town3TalkStart, Town3TalkEnd)
	ns.SetDialog(warden, ns.NORMAL, Jail1Start, Jail1End)
}
func PlayIxTownMusic() {
	if !ns.IsCaller(ns.GetHost()) {
		goto LABEL1
	}
	ns.Music(15, 40)
LABEL1:
	return
}
func DisableElevators() {
	PlayIxTownMusic()
	ns.JournalEntry(ns.GetHost(), "LocateAldwyn", 2)
	ns.PrintToAll("Con01a:NewJournalEntry")
	ns.ObjectOff(aldwynElev)
	ns.ObjectOff(mayorElev)
	ns.ObjectOff(magicShopElev)
}
func PickupObjects() {
	ns.Pickup(contestOfficial, contestPrize)
	ns.Pickup(contestOfficial, contestQuiver1)
	ns.Pickup(contestOfficial, contestQuiver2)
	ns.Pickup(morgan, cheapBow)
}
func DebugOff() {
	flag118 = false
	ns.PrintToAll("Debug mode has been disabled.")
}
func AldwinCheck() {
	if !ns.HasItem(aldwyn, charmSpell) {
		goto LABEL1
	}
	ns.PrintToAll("Spell is in Aldwins inv.")
	goto LABEL2
LABEL1:
	ns.PrintToAll("Spell not found in Aldwins inv.")
LABEL2:
	if !ns.HasItem(aldwyn, ratFieldGuide) {
		goto LABEL3
	}
	ns.PrintToAll("Field Guide is in Aldwins inv.")
	goto LABEL4
LABEL3:
	ns.PrintToAll("Field Guide not found in Aldwins inv.")
LABEL4:
	return
}
func PlayUrchinDenMusic() {
	ns.Music(17, 100)
}
func PlayTunnelMusic() {
	ns.Music(18, 100)
}
func ExitDen() {
	gvar211 = gvar205
	ns.SetDialog(warden, ns.NORMAL, Jail2Start, Jail2End)
	ns.MoveObject(obj123, 2969, 2354)
	ns.MoveObject(obj124, 3016, 2308)
	ns.CreatureGuard(obj123, 2969, 2354, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.CreatureGuard(obj124, 3016, 2308, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetOwner(ns.GetHost(), obj123)
	ns.SetOwner(ns.GetHost(), obj124)
}
func Town3TalkStart() {
	ns.Frozen(tommy, true)
	ns.CreatureIdle(tommy)
	ns.LookAtObject(tommy, ns.GetHost())
	if gvar210 != gvar202 {
		goto LABEL1
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman9Talk01")
	return
LABEL1:
	if gvar210 == gvar207 {
		goto LABEL3
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman9Talk02")
LABEL3:
	if gvar210 != gvar207 {
		return
	}
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman9Talk03")
	return
}
func Town3TalkEnd() {
	ns.Frozen(tommy, false)
	ns.Wander(tommy)
	ns.SetDialog(tommy, ns.NORMAL, Town3TalkStart, Town3TalkEnd)
}
func Maiden2TalkStart() {
	ns.Frozen(gretchen, true)
	ns.CreatureIdle(gretchen)
	ns.LookAtObject(gretchen, ns.GetHost())
	if gvar211 == gvar205 {
		ns.TellStory(ns.HumanMaleEatFood, "Con02A:Maiden2Talk01")
	} else {
		ns.TellStory(ns.HumanMaleEatFood, "Con02A:Townsman2Talk02")
	}
}
func Maiden2TalkEnd() {
	ns.Frozen(gretchen, false)
	ns.Wander(gretchen)
	ns.SetDialog(gretchen, ns.NORMAL, Maiden2TalkStart, Maiden2TalkEnd)
}
func HenrickGreetStart() {
	ns.DestroyEveryChat()
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk03")
}
func HenrickGreetEnd() {
	ns.SetDialog(henrick, ns.NORMAL, HenrickStart, HenrickEnd)
}
func HenrickStart() {
	ns.DestroyEveryChat()
	if gvar213 == gvar202 {
		switch ns.Random(1, 3) {
		case 1:
			ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk04")
		case 2:
			ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk05")
		case 3:
			ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk03")
		}
	}
	if gvar213 != gvar204 {
		return
	}
	switch ns.Random(1, 2) {
	case 1:
		ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk04")
	case 2:
		ns.TellStory(ns.HumanMaleEatFood, "Con02A:HenrickTalk06")
	}
}
func HenrickEnd() {
	ns.SetDialog(henrick, ns.NORMAL, HenrickStart, HenrickEnd)
}
func Jail1Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk01")
}
func Jail1End() {
	ns.SetDialog(warden, ns.NORMAL, Jail1Start, Jail1End)
}
func Jail2Start() {
	ns.TellStory(ns.HumanMaleEatFood, "Con02A:JailerTalk02")
}
func Jail2End() {
	ns.SetDialog(warden, ns.NORMAL, Jail2Start, Jail2End)
}
func MakeAWish() {
	v0 := ns.MaxHealth(ns.GetCaller())
	ns.RestoreHealth(ns.GetCaller(), v0-ns.CurrentHealth(ns.GetCaller()))
	ns.PrintToAll("GeneralPrint:WellSignRefresh")
	ns.AudioEvent(ns.RestoreHealthName, wp193)
}
func SetupCreatures() {
	ns.SetOwner(ns.GetHost(), henrick)
	ns.SetOwner(ns.GetHost(), warden)
	ns.SetOwner(ns.GetHost(), pepper)
	ns.LockDoor(necroDoor)
	ns.LockDoor(obj214)
	ns.LockDoor(obj215)
	ns.LockDoor(obj216)
	ns.LockDoor(obj217)
}
func StartWolfAttack() {
	if flag220 {
		return
	}
	ns.ObjectGroupOff(gvar218)
	ns.WideScreen(true)
	ns.Frozen(ns.GetHost(), true)
	ns.CreatureIdle(ns.GetHost())
	ns.LookAtObject(ns.GetHost(), henrick)
	ns.MoveObject(pepper, ns.GetWaypointX(wp219), ns.GetWaypointY(wp219))
	ns.FrameTimer(1, ActivateWolf)
}
func ActivateWolf() {
	ns.CreatureIdle(pepper)
	ns.AggressionLevel(pepper, 0.83)
	ns.Attack(pepper, ns.GetHost())
	flag220 = true
	ns.FrameTimer(50, HenrickWarning)
}
func HenrickWarning() {
	ns.Chat(henrick, "Con02A:HenrickTalk01")
	ns.Frozen(henrick, true)
	ns.CreatureIdle(henrick)
	ns.LookAtObject(henrick, pepper)
	ns.Frozen(pepper, true)
	ns.LookAtObject(pepper, henrick)
	ns.MoveWaypoint(wp113, ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()))
	ns.AudioEvent(ns.CharmCast, wp113)
	ns.FrameTimer(20, CharmPepper)
}
func CharmPepper() {
	ns.Effect(ns.CHARM, ns.GetObjectX(henrick), ns.GetObjectY(henrick), ns.GetObjectX(pepper), ns.GetObjectY(pepper))
	ivar221++
	if !(ivar221 < 60) {
		CharmEnd()
		return
	}
	ns.FrameTimer(1, CharmPepper)
	return
}
func CharmEnd() {
	ns.AudioEvent(ns.CharmSuccess, wp113)
	ns.WideScreen(false)
	ns.Frozen(ns.GetHost(), false)
	ns.Frozen(henrick, false)
	ns.Frozen(pepper, false)
	ns.CreatureGuard(henrick, ns.GetObjectX(henrick), ns.GetObjectY(henrick), ns.GetObjectX(ns.GetHost()), ns.GetObjectY(ns.GetHost()), 0)
	ns.SetOwner(ns.GetHost(), pepper)
	ns.AggressionLevel(pepper, 0.16)
	ns.CreatureFollow(pepper, henrick)
	ns.Chat(henrick, "Con02A:HenrickTalk02")
	ns.SetDialog(henrick, ns.NORMAL, HenrickGreetStart, HenrickGreetEnd)
}
func OnEvent(typ string) {
	switch typ {
	case "PlayerDeath":
		PlayerDeath()
	case "MapInitialize":
		MapInitialize()
	}
}
