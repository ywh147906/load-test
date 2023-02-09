package serverids

import (
	"strconv"

	"github.com/ywh147906/load-test/common/consulkv"
	"github.com/ywh147906/load-test/common/utils"
	wr "github.com/ywh147906/load-test/common/utils/weightedrand"
	"github.com/ywh147906/load-test/common/values"
)

type serverIds struct {
	chooser *wr.Chooser[values.ServerId, int]
}

var sids *serverIds

func Init(cnf *consulkv.Config) {
	ids := make(map[string]int)
	err := cnf.Unmarshal("server_ids", &ids)
	utils.Must(err)

	choices := make([]*wr.Choice[values.ServerId, int], 0, len(ids))
	for sid, weight := range ids {
		id, err := strconv.Atoi(sid)
		utils.Must(err)
		choices = append(choices, wr.NewChoice(values.ServerId(id), weight))
	}

	ch, err := wr.NewChooser(choices...)
	utils.Must(err)
	sids = &serverIds{chooser: ch}
}

func (s *serverIds) Assign() values.ServerId {
	return s.chooser.Pick()
}

func Assign() values.ServerId {
	return sids.Assign()
}
