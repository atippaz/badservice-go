package service

import (
	_matchModel "bad-service-go/pkg/match/model"
	_matchRepository "bad-service-go/pkg/match/repository"
	"fmt"
	"strings"
	"time"
)

type matchServiceImpl struct {
	matchRepository _matchRepository.MatchRepository
}

func NewmatchServiceImpl(matchRepository _matchRepository.MatchRepository) MatchService {
	return &matchServiceImpl{
		matchRepository: matchRepository,
	}
}

func (r *matchServiceImpl) DeleteById(matchId string) ( error) {
	err := r.matchRepository.DeleteById(matchId)
	if err != nil {
		return err
	}
	return nil
}

func (r *matchServiceImpl) Insert(matchRequest *_matchModel.MatchInsertRequest) (*_matchModel.MatchResult, error) {
	members := matchRequest.Members;
	if members == nil {
		*members = make([]string, 0) 
	}
	limit := matchRequest.Limit
	if limit == nil {
		*limit = 2
	}
	start := matchRequest.Start
	if start == nil {
		*start = 0
	}
	random := matchRequest.Random
	if random == nil {
		*random = false
	}
	isSet :=	matchRequest.MatchSet.IsSet
	limitSet :=	matchRequest.MatchSet.LimitSet
	// teamLock :=	matchRequest.teamLock
	if strings.TrimSpace(matchRequest.SetName) == "" {
		// โหลด time zone ของกรุงเทพฯ
		location, err := time.LoadLocation("Asia/Bangkok")
		if err == nil {
			now := time.Now().In(location)
			formattedDate := fmt.Sprintf("วันที่ %02d-%02d-%04d %02d:%02d",
			now.Day(), int(now.Month()), now.Year(),
			now.Hour(), now.Minute())
			matchRequest.SetName = formattedDate
		}
	}
	// if (isSet) {
	// 	const teams = await createTeamPair(
	// 	  members,
	// 	  limit,
	// 	  limitSet,
	// 	  start,
	// 	  random,
	// 	  req.body.teamLock.map((x) => x.teamMember)
	// 	);
  
	// 	const payload = {
	// 	  courtNumber: req.body.courtNumber,
	// 	  roomId: req.body.roomId,
	// 	  winScore: req.body.winScore,
	// 	  teamLimit: req.body.teamLimit,
	// 	  winStreak: req.body.winStreak,
	// 	  allTeam: teams.map((x, index) => {
	// 		return {
	// 		  order: index + 1,
	// 		  set: x.map((y, idx) => {
	// 			return {
	// 			  order: idx + 1,
	// 			  member: y,
	// 			};
	// 		  }),
	// 		};
	// 	  }),
	// 	  setName: req.body.setName,
	// 	};
	// 	const team = new MatchSetDB(payload);
	// 	await team.save().then((e) => res.json({ isSet: true, id: e._id }));
	//   } else {
	// 	const allteam = createRandomTeamParing(members, teamLock, limit);
	// 	const payload = {
	// 	  courtNumber: req.body.courtNumber,
	// 	  roomId: req.body.roomId,
	// 	  winScore: req.body.winScore,
	// 	  teamLimit: req.body.teamLimit,
	// 	  winStreak: req.body.winStreak,
	// 	  allTeam: allteam.map((x, index) => {
	// 		return { member: x, order: index + 1 };
	// 	  }),
	// 	  setName: req.body.setName,
	// 	};
	// 	return await MatchDB(payload)
	// 	  .save()
	// 	  .then((e) => res.json({ isSet: false, id: e._id }));
	//   }


	// newmatch := entities.Match{
	// 	// Email:  matchRequest.Email,
	// 	// Avatar: matchRequest.Avatar,
	// 	// Name:   matchRequest.Name,
	// }
	// _, err := r.matchRepository.Insert(newmatch)
	// if err != nil {
	// 	return nil, err
	// }
	return &_matchModel.MatchResult{
		// ID:     result.ID,
		// Email:  result.Email,
		// Name:   result.Name,
		// Avatar: result.Avatar,
	}, nil
}

func (r *matchServiceImpl) FindAll() (*[]_matchModel.MatchResult, error) {
	result, error := r.matchRepository.FindAll()
	if error != nil {
		return nil, error
	}
	var res []_matchModel.MatchResult
	for _,data  := range result{
		res = append(res, _matchModel.MatchResult{
		RoomId: data.RoomID,
		TeamId: data.ID.Hex(),
		TeamName: data.SetName,
		CourtNumber:data.CourtNumber,
		AllTeam: data.AllTeam,
		WinScore: data.WinScore,
		TeamLimit: data.TeamLimit,
		WinStreak: data.WinStreak,
		})
	}
	return &res, nil
}

func (r *matchServiceImpl) FindById(matchId string) (*_matchModel.MatchResult, error) {
	data, error := r.matchRepository.FindById(matchId)
	if error != nil {
		return nil, error
	}
	return &_matchModel.MatchResult{
		RoomId: data.RoomID,
		TeamId: data.ID.Hex(),
		TeamName: data.SetName,
		CourtNumber:data.CourtNumber,
		AllTeam: data.AllTeam,
		WinScore: data.WinScore,
		TeamLimit: data.TeamLimit,
		WinStreak: data.WinStreak,
		}, nil
}


  func shufferMember(member) {
	for (let i = member.length - 1; i > 0; i--) {
	  const j = Math.floor(Math.random() * (i + 1));
	  [member[i], member[j]] = [member[j], member[i]];
	}
	return member;
  }
  func createRandomTeamParing(member, teamLock, limit = 2) {
	const result = [];
	member.forEach((x) => {
	  if (result.length == 0) {
		result.push([x]);
	  } else if (result[result.length - 1].length < limit) {
		result[result.length - 1].push(x);
	  } else {
		result.push([x]);
	  }
	});
	teamLock.forEach((x) => {
	  result.push(x.teamMember);
	});
	return shufferMember(result);
  }
  
  func findOldMember(e, _histories) {
	const oldMember = [];
	const histories = JSON.parse(JSON.stringify(_histories));
	histories.forEach((h) => {
	  const filter = h.filter((x) => x.some((s) => s == e));
	  if (filter.length <= 0) {
		return;
	  }
	  const paired = filter.flatMap((x) => x).filter((x) => x != e);
	  if (paired.length >= 1) {
		oldMember.push(paired);
	  }
	});
	return oldMember;
  }