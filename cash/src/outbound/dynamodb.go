package outbound

import (
	"context"
	"fmt"

	"payment-cash/src/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type paymentStruct struct {
	db *dynamo.DB
}

func NewPaymentDB() domain.PaymentRepository {
	db, err := createSession()
	if err != nil {
		panic(err)
	}
	return &paymentStruct{db: db}
}

// 決済データの登録
func (p *paymentStruct) PutItem(ctx context.Context, model *domain.Payment) (*domain.Payment, error) {
	paymentNo, err := sequence(ctx, p.db)
	if err != nil {
		return nil, err
	}
	model.PaymentNo = fmt.Sprintf("P%10d", *paymentNo) // dummy 支払い番号はP始まりとする
	model.DeleteFlag = false
	table := p.db.Table("Payment")
	if err := table.Put(model).Run(); err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
		return nil, err
	}
	fmt.Println("accept : " + model.AcceptNo)
	return model, nil
}

// OrderNoに該当する決済データを取得
func (p *paymentStruct) GetItem(ctx context.Context, orderNo string) (*domain.Payment, error) {
	table := p.db.Table("Payment")
	var result domain.Payment
	err := table.Get("OrderNo", orderNo).One(&result)
	if err != nil {
		return nil, err
	}
	if result.DeleteFlag {
		return nil, nil
	}
	return &result, nil
}

// orderNoに該当する決済データ論理を削除
func (p *paymentStruct) DeleteItem(ctx context.Context, orderNo string) error {
	table := p.db.Table("Payment")

	var result domain.Payment
	return table.Update("OrderNo", orderNo).Set("DeleteFlag", true).Value(&result)
}

func createSession() (*dynamo.DB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://payment-cash-localstack.nautible-app-ms.svc.cluster.local:4566"),
		Credentials: credentials.NewStaticCredentials("test-key", "test-secret", ""),
	})
	if err != nil {
		return nil, err
	}
	db := dynamo.New(sess)
	return db, nil
}

// シーケンス取得
func sequence(ctx context.Context, db *dynamo.DB) (*int, error) {
	var counter struct {
		Name           string
		SequenceNumber int
	}
	table := db.Table("Sequence")
	err := table.Update("Name", "Payment").Add("SequenceNumber", 1).ValueWithContext(ctx, &counter)
	if err != nil {
		return nil, err
	}
	return &counter.SequenceNumber, err

}
