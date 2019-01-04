package datastore

import (
	"context"
	"os"

	pb "github.com/micnncim/mercari-datastore-sample/proto"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/boom"
	"go.mercari.io/datastore/clouddatastore"
)

type Client struct {
	bm *boom.Boom
}

func FromContext(ctx context.Context) (*Client, error) {
	ds, err := clouddatastore.FromContext(ctx, datastore.WithProjectID(os.Getenv("PROJECT_ID")))
	if err != nil {
		return nil, err
	}
	bm := boom.FromClient(ctx, ds)
	return &Client{
		bm: bm,
	}, nil
}

func (c *Client) CreateUser(ctx context.Context, u *pb.User) error {
	if _, err := c.bm.Put(u); err != nil {
		return err
	}
	return nil
}

func (c *Client) ListUsers(ctx context.Context) ([]*pb.User, error) {
	var us []*pb.User
	q := c.bm.NewQuery(c.bm.Kind(us))
	if _, err := c.bm.GetAll(q, &us); err != nil {
		return nil, err
	}
	return us, nil
}

func (c *Client) UpdateUser(ctx context.Context, u *pb.User) (*pb.User, error) {
	_, err := c.bm.RunInTransaction(func(tx *boom.Transaction) error {
		uu := &pb.User{
			ID: u.ID,
		}
		if err := c.bm.Get(uu); err != nil {
			return err
		}
		if _, err := c.bm.Put(u); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *Client) DeleteUser(ctx context.Context, u *pb.User) error {
	return c.bm.Delete(u)
}
