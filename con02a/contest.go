package con02a

import "github.com/noxworld-dev/noxscript/ns/v3"

var (
	flag4         bool
	contestLastWP int
	ivar7         int
	ivar8         int
	contestDelay  = 85
	ivar61        int
	flag64        bool

	contestOfficial ns.ObjectID
	contestGuard    ns.ObjectID
	contestBarrel   ns.ObjectID
	heckler         ns.ObjectID

	contestGate1 ns.ObjectID
	contestGate2 ns.ObjectID
	contestGate3 ns.ObjectID

	contestQuiver1 ns.ObjectID
	contestQuiver2 ns.ObjectID
	contestPrize   ns.ObjectID
)

func ContestWon(n int) {
	ns.LookAtObject(contestOfficial, ns.GetHost())
	flag64 = true
	gvar54 = gvar34
	ns.StartDialog(contestOfficial, ns.GetHost())
	ns.UnlockDoor(contestGate1)
	ns.UnlockDoor(contestGate2)
}

func ContestLost(n int) {
	ns.DestroyEveryChat()
	ivar61++
	ns.LookAtObject(contestOfficial, ns.GetHost())
	if ivar61 == 1 {
		gvar54 = gvar32
	}
	if ivar61 > 1 {
		ns.SetDialog(contestOfficial, ns.NORMAL, ContestOfficialDialogStart, ContestOfficialDialogEnd)
		ns.SetShopkeeperText(barkeeper, "Con02:BarkeeperDefault")
		gvar54 = gvar35
	}
	switch n {
	case 8:
		ns.Chat(heckler, "Con02a:Contest8of10")
	case 7:
		ns.Chat(heckler, "Con02a:Contest7of10")
	case 6:
		ns.Chat(heckler, "Con02a:Contest6of10")
	case 5:
		ns.Chat(heckler, "Con02a:Contest5of10")
	case 4:
		ns.Chat(heckler, "Con02a:Contest4of10")
	case 3:
		ns.Chat(heckler, "Con02a:Contest3of10")
	case 2:
		ns.Chat(heckler, "Con02a:Contest2of10")
	case 1:
		ns.Chat(heckler, "Con02a:Contest1of10")
	case 0:
		ns.Chat(heckler, "Con02a:Contest0of10")
	}
	ns.UnlockDoor(contestGate1)
	ns.UnlockDoor(contestGate2)
}

func WinContest() {
	gvar210 = gvar207
}

func ContestEngine() {
	var wpi int
	hp := ns.CurrentHealth(contestBarrel)
	if flag4 {
		wpi = 1
		flag4 = false
	} else {
		wpi = ns.Random(0, 2)
		for {
			if wpi != contestLastWP {
				break
			}
			wpi = ns.Random(0, 2)
		}
		contestLastWP = wpi
		if hp <= 0 {
			ivar7++
			contestDelay = contestDelay - ivar10
		} else {
			ns.Delete(contestBarrel)
			ivar8++
		}
	}
	if ivar8+ivar7 >= 10 {
		if ivar7 >= 9 && ivar8 <= 1 {
			ContestWon(ivar7)
			WinContest()
			ivar7 = 0
			ivar8 = 0
			contestLastWP = 1
			contestDelay = 60
			flag4 = true
		} else {
			ContestLost(ivar7)
			ivar7 = 0
			ivar8 = 0
			contestLastWP = 1
			contestDelay = 85
			flag4 = true
		}
		return
	}
	switch wpi {
	case 0:
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp13), ns.GetWaypointY(wp13), 0, 0)
		contestBarrel = ns.CreateObject("TargetBarrel1", wp13)
		ns.AudioEvent(ns.BlinkCast, wp16)
	case 1:
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp14), ns.GetWaypointY(wp14), 0, 0)
		contestBarrel = ns.CreateObject("TargetBarrel1", wp14)
		ns.AudioEvent(ns.BlinkCast, wp16)
	case 2:
		ns.Effect(ns.SMOKE_BLAST, ns.GetWaypointX(wp15), ns.GetWaypointY(wp15), 0, 0)
		contestBarrel = ns.CreateObject("TargetBarrel1", wp15)
		ns.AudioEvent(ns.BlinkCast, wp16)
	}
	ns.LookAtObject(contestOfficial, contestBarrel)
	ns.FrameTimer(contestDelay, ContestEngine)
}

func ResetContestGuardDialogChoice() {
	ns.SetDialog(contestGuard, ns.YESNO, ContestGuardDialogStart, ContestGuardDialogEnd)
}

func ResetContestOfficialDialogChoice() {
	ns.SetDialog(contestOfficial, ns.YESNO, ContestOfficialDialogStart, ContestOfficialDialogEnd)
}

func ContestOfficialDialogStart() {
	ns.DestroyEveryChat()
	switch gvar54 {
	case gvar32:
		ns.LookAtObject(contestOfficial, ns.GetHost())
		if !(ivar61 > 0) {
			ns.TellStory(ns.SwordsmanHurt, "Con02a:ContestGreeting")
		} else {
			ns.SetDialog(contestOfficial, ns.YESNO, NullDialogStart, ContestOfficialDialogEnd)
			ns.TellStory(ns.SwordsmanHurt, "Con02a:RetryContestChoice")
		}
	case gvar33:
		ns.LookAtObject(contestOfficial, ns.GetHost())
		ns.TellStory(ns.SwordsmanHurt, "Con02a:ContestOfficialWaiting")
	case gvar34:
		ns.LookAtObject(contestOfficial, ns.GetHost())
		ns.TellStory(ns.SwordsmanHurt, "Con02a:AwardContestPrize")
		ns.SetShopkeeperText(barkeeper, "Con02:BarkeeperDefault")
	case gvar35:
		ns.LookAtObject(contestOfficial, ns.GetHost())
		ns.TellStory(ns.SwordsmanHurt, "Con02a:NoMoreContestTries")
	}
}

func ContestOfficialDialogEnd() {
	answer := ns.GetAnswer(contestOfficial)
	gold := ns.GetGold(ns.GetHost())
	switch gvar54 {
	case gvar32:
		if !(ivar61 > 0) {
			ns.Pickup(ns.GetHost(), contestQuiver1)
			ns.PrintToAll("Con02a:QuiverAddedToInventory")
			EnableFiringLineTrigger()
			ns.LockDoor(contestGate1)
			ns.LockDoor(contestGate2)
			gvar54 = gvar33
			return
		}
		if answer != 1 {
			ns.SetDialog(contestOfficial, "FALSE", ContestOfficialDialogStart, ContestOfficialDialogEnd)
			return
		}
		if !(gold >= contestBowCost) {
			ns.SetDialog(contestOfficial, ns.NORMAL, NullDialogStart, ResetContestOfficialDialogChoice)
			ns.TellStory(ns.SwordsmanHurt, "Con02a:NotEnoughGold")
			return
		}
		ns.ChangeGold(ns.GetHost(), -contestBowCost)
		ns.SetDialog(contestOfficial, ns.NORMAL, ContestOfficialDialogStart, EnableFiringLineTrigger)
		ns.TellStory(ns.SwordsmanHurt, "Con02a:RetryContest")
		ns.Pickup(ns.GetHost(), contestQuiver2)
		ns.PrintToAll("Con02a:QuiverAddedToInventory")
		ns.LockDoor(contestGate1)
		ns.LockDoor(contestGate2)
		gvar54 = gvar33
	case gvar33:
		// noop
	case gvar34:
		ns.Pickup(ns.GetHost(), contestPrize)
		gvar54 = gvar35
		ns.FrameTimer(15, AwardContestExperience)
	case gvar35:
		// noop
	}
}

func ContestGuardDialogStart() {
	var v0 int
	v0 = gvar51
	if v0 == gvar22 {
		goto LABEL1
	}
	if v0 == gvar23 {
		goto LABEL2
	}
	if v0 == gvar24 {
		goto LABEL3
	}
	goto LABEL4
LABEL1:
	ns.LookAtObject(contestGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:ContestGuard")
	goto LABEL4
LABEL2:
	ns.LookAtObject(contestGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:EnterContest")
	flag63 = true
	ns.UnlockDoor(contestGate1)
	ns.UnlockDoor(contestGate2)
	goto LABEL4
LABEL3:
	ns.LookAtObject(contestGuard, ns.GetHost())
	ns.TellStory(ns.SwordsmanHurt, "Con02a:RandomSay")
	goto LABEL4
LABEL4:
	return
}

func ContestGuardDialogEnd() {
	var (
		v0 int
		v1 int
		v2 int
		v3 int
		v4 int
		v5 bool
		v6 ns.ObjectID
		v7 int
		v8 int
	)
	v0 = 0
	v1 = 0
	v2 = 1
	v3 = 2
	v4 = 0
	v5 = false
	v6 = ns.GetLastItem(ns.GetHost())
	for {
		if v6 == 0 {
			goto LABEL1
		}
		if !(ns.HasClass(v6, ns.WEAPON) && ns.HasSubclass(v6, "BOW")) {
			goto LABEL2
		}
		v5 = true
		goto LABEL1
	LABEL2:
		v6 = ns.GetPreviousItem(v6)
	}
LABEL1:
	v8 = gvar51
	if v8 == gvar22 {
		goto LABEL4
	}
	if v8 == gvar23 {
		goto LABEL5
	}
	if v8 == gvar24 {
		goto LABEL6
	}
	goto LABEL7
LABEL4:
	v1 = ns.GetAnswer(contestGuard)
	v7 = v1
	if v7 == v2 {
		goto LABEL8
	}
	if v7 == v3 {
		goto LABEL9
	}
	if v7 == v4 {
		goto LABEL10
	}
	goto LABEL11
LABEL8:
	v0 = ns.GetGold(ns.GetHost())
	if !(v0 < contestBowCost) {
		goto LABEL12
	}
	ns.SetDialog(contestGuard, ns.NORMAL, NullDialogStart, ResetContestGuardDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:NotEnoughGold")
	goto LABEL13
LABEL12:
	if v5 {
		goto LABEL14
	}
	ns.SetDialog(contestGuard, ns.NORMAL, NullDialogStart, ResetContestGuardDialogChoice)
	ns.TellStory(ns.SwordsmanHurt, "Con02a:NoBow")
	goto LABEL13
LABEL14:
	ns.ChangeGold(ns.GetHost(), -contestBowCost)
	gvar51 = gvar23
	ns.SetDialog(contestGuard, ns.NORMAL, ContestGuardDialogStart, ContestGuardDialogEnd)
	ContestGuardDialogStart()
LABEL13:
	goto LABEL11
LABEL9:
	goto LABEL11
LABEL10:
	goto LABEL11
LABEL11:
	goto LABEL7
LABEL5:
	ns.CancelDialog(contestGuard)
	goto LABEL7
LABEL6:
	goto LABEL7
LABEL7:
	return
}
