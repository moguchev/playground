package main

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	tr "gitlab.services.mts.ru/pepperpotts/api/oebs/transferservice"
	"google.golang.org/grpc"
)

var oebsCli tr.TransferOEBSClient

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:9081", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		conn.Close()
	}()

	oebsCli := tr.NewTransferOEBSClient(conn)
	ctx := context.Background()
	///////////////////////////////////////////////////////////////////////////////////////////

	// ЗП
	sal, err := oebsCli.GetSalaryByAssignments(ctx, &tr.GetSalaryByAssignmentsRequest{
		AssignmentId: []int64{355716},
		UseCache:     false,
	})
	if err != nil {
		fmt.Println(fmt.Errorf("ЗП: %w", err))
	} else {
		dateFrom, err := ptypes.Timestamp(sal.Info[0].DateFrom)
		if err != nil {
			fmt.Println(fmt.Errorf("ЗП: %w", err))
		}
		fmt.Println("GetSalaryByAssignments")
		fmt.Println("date_from:", dateFrom.Format("02.01.2006"))
	}

	// Карточка ШЕ
	card, err := oebsCli.GetPositionForCard(ctx, &tr.GetPositionForCardRequest{
		Id:            2943179,
		AssignmentsId: []int64{355716},
	})
	if err != nil {
		fmt.Println(fmt.Errorf("Карточка ШЕ: %w", err))
	} else {
		ass := card.Assignments[0].Assignments
		fmt.Println("GetPositionForCard")
		for i := range ass {
			startDate, err := ptypes.Timestamp(ass[i].StartDate)
			if err == nil {
				fmt.Println("start_date: ", startDate.Format("02.01.2006"))
			}
		}
	}

	// Бюджет ШЕ
	// fot, err := oebsCli.GetFotByPositions(ctx, &tr.GetFotByPositionsRequest{
	// 	PositionId: []int64{202080},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Бюджет ШЕ: %w", err))
	// } else {
	// 	fmt.Println(*fot)
	// }

	// Бонусы
	// bon, err := oebsCli.GetBonusesByAssignments(ctx, &tr.GetBonusesByAssignmentsRequest{
	// 	AssignmentId: []int64{116290},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Бонусы: %w", err))
	// } else {
	// 	fmt.Println(bon.Recv())
	// }

	// информация для валидации
	// info, err := oebsCli.GetValidationInfo(ctx, &tr.GetValidationInfoRequest{
	// 	AdditionalPlace:      "",
	// 	AssignmentObjVersion: nil,
	// 	CurrentAssignmentId:  116290,
	// 	PositionObjVersion:   nil,
	// 	TargetPositionId:     202080,
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("информация для валидации: %w", err))
	// } else {
	// 	fmt.Println(*info)
	// }

	// EP
	// ep, err := oebsCli.GetEarningPolicy(ctx, &tr.GetEarningPolicyRequest{
	// 	RegionId: 0,
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("EP: %w", err))
	// } else {
	// 	fmt.Println(ep.Recv())
	// }

	// EP info
	// epinfo, err := oebsCli.GetExtraInfoByEarningPolicyID(ctx, &tr.GetExtraInfoByEarningPolicyIDRequest{
	// 	EarningPolicyId: 11478,
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("EP info: %w", err))
	// } else {
	// 	fmt.Println(*epinfo)
	// }

	// schedules
	// ss, err := oebsCli.GetSchedules(ctx, &tr.GetSchedulesRequest{
	// 	EarningPolicyId: 11478,
	// 	Limit:           1,
	// 	Name:            "nil",
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("schedules: %w", err))
	// } else {
	// 	fmt.Println(ss.Recv())
	// }

	// Условия работы
	// wc, err := oebsCli.GetWorkConditions(ctx, &tr.GetWorkConditionsRequest{
	// 	AssignmentId: 116290,
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Условия работы: %w", err))
	// } else {
	// 	fmt.Println(wc)
	// }

	// Прошлые назначения
	// prev, err := oebsCli.GetPreviousAssignment(ctx, &tr.GetPreviousAssignmentRequest{
	// 	Params: []*tr.PreviousRequestParameters{
	// 		&tr.PreviousRequestParameters{
	// 			AssignmentId: 116290,
	// 			PositionId:   202080,
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Прошлые назначения: %w", err))
	// } else {
	// 	fmt.Println(prev.Recv())
	// }

	///////////////////////////////////////////////////////////////////
	//!!!
	// Прошлые условия труда
	// pwc, err := oebsCli.GetPreviousWorkConditions(ctx, &tr.GetPreviousWorkConditionsRequest{
	// 	Params: &tr.PreviousRequestParameters{
	// 		AssignmentId: 116290,
	// 		PositionId:   116290,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Прошлые условия труда: %w", err))
	// } else {
	// 	fmt.Println(*pwc)
	// }
	///////////////////////////////////////////////////////////////////////////

	// Проверка условий труда
	// b, err := oebsCli.IsWorkConditionCorrect(ctx, &tr.IsWorkConditionCorrectRequest{
	// 	EarningPolicyId: 11478,
	// 	IsNormDay:       true,
	// 	PayBasisId:      64,
	// 	PlanId:          1,
	// 	WorkModeCode:    63,
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Проверка условий труда: %w", err))
	// } else {
	// 	fmt.Println(*b)
	// }

	// Карточка ШЕ
	// card, err := oebsCli.GetPositionForCard(ctx, &tr.GetPositionForCardRequest{
	// 	Id:            2926286,
	// 	AssignmentsId: []int64{116290},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Карточка ШЕ: %w", err))
	// } else {
	// 	fmt.Println(*card)
	// }

	// // ЕС описание
	// ecd, err := oebsCli.GetEcDescription(ctx, &tr.GetEcDescriptionRequest{
	// 	Code: []string{"11"},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("ЕС описание: %w", err))
	// } else {
	// 	fmt.Println(*ecd)
	// }

	// Назначения на ШЕ
	// as, err := oebsCli.GetAssignmentsForPosition(ctx, &tr.GetAssignmentsForPositionRequest{
	// 	PositionId: 202080,
	// 	Ids:        []int64{202080},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Назначения на ШЕ: %w", err))
	// } else {
	// 	fmt.Println(*as)
	// }

	// Проверка шедула
	// cor, err := oebsCli.IsScheduleCorrect(ctx, &tr.IsScheduleCorrectRequest{
	// 	AssignmentId: 116290,
	// 	PlanId:       1,
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Проверка шедула: %w", err))
	// } else {
	// 	fmt.Println(*cor)
	// }

	// Инфо по ШЕ для HR
	// i, err := oebsCli.GetInfoForHRDbyPositions(ctx, &tr.GetInfoForHRDbyPositionsRequest{
	// 	PositionId: []string{"116290", "202080"},
	// })
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("Инфо по ШЕ для HR: %w", err))
	// } else {
	// 	fmt.Println(*i)
	// }

	// MakeTransfer(ctx context.Context, in *MakeTransferRequest, opts ...grpc.CallOption) (*MakeTransferResponse, error)
	// MakeChangeRoration(ctx context.Context, in *MakeChangeRotationRequest, opts ...grpc.CallOption) (*MakeChangeRotationResponse, error)
}
