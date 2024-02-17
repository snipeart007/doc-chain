package server

import (
	"context"
	"time"

	"github.com/snipeart007/doc-chain/base_blockchain/chain"
	pb "github.com/snipeart007/doc-chain/bootstrap/data_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *DataServiceServer) GetDocument(ctx context.Context, id *pb.DocumentID) (*pb.Document, error) {
	document, err := server.blockchain.RetreiveDocument(id.Id)
	if err == chain.ErrDocumentDoesNotExist {
		return nil, status.Error(codes.NotFound, "specified document does not exist")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Document{
		Id:        document.ID,
		Data:      document.Data,
		Timestamp: timestamppb.New(time.Unix(document.TimeStamp, 0)),
		Block:     document.Block,
	}, nil
}

func (server *DataServiceServer) InsertDocument(ctx context.Context, args *pb.CreateDocumentArgs) (*pb.DocumentID, error) {
	id, err := server.blockchain.InsertJSONDocument(args.Data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DocumentID{
		Id: id,
	}, nil
}

func (server *DataServiceServer) StartMining(stream pb.DataService_StartMiningServer) error {
	for {
		// in, err := stream.Recv()
		// if err == io.EOF {
		// 	return nil
		// }
		// if err != nil {
		// 	return err
		// }
		// key := 0
		// for _, note := range s.routeNotes[key] {
		// 	if err := stream.Send(note); err != nil {
		// 		return err
		// 	}
		// }
		
	}
}
