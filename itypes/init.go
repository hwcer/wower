package itypes

import (
	"errors"
	"github.com/hwcer/updater"
	"github.com/hwcer/updater/operator"
	"github.com/hwcer/wower/model"
)

type Unique func(u *updater.Updater, iid int32) (string, error)
type Creator func(u *updater.Updater, iid int32, val int64) (r any, err error)
type Listener func(u *updater.Updater, op *operator.Operator)

func NewIType(id int32, l ...Listener) *IType {
	it := &IType{id: id, stacked: true}
	if len(l) > 0 {
		it.listener = l[0]
	}
	return it
}

type IType struct {
	id       int32
	unique   Unique
	stacked  bool
	creator  Creator
	listener Listener
}

func (this *IType) ID() int32 {
	return this.id
}

func (this *IType) New(u *updater.Updater, op *operator.Operator) (any, error) {
	return this.Create(u, op.IID, op.Value)
}

func (this *IType) Create(u *updater.Updater, iid int32, val int64) (any, error) {
	if this.creator != nil {
		return this.creator(u, iid, val)
	} else {
		return nil, errors.New("create fail")
	}
}

func (this *IType) Stacked() bool {
	return this.stacked
}

func (this *IType) ObjectId(u *updater.Updater, iid int32) (string, error) {
	if this.unique != nil {
		return this.unique(u, iid)
	}
	if this.stacked {
		return model.Unique(u, iid)
	} else {
		return model.ObjectId.New(uint32(iid)), nil
	}
}

func (this *IType) Listener(u *updater.Updater, op *operator.Operator) {
	if this.listener != nil {
		this.listener(u, op)
	}
}

func (this *IType) SetUnique(unique Unique) {
	this.unique = unique
}
func (this *IType) SetStacked(stacked bool) {
	this.stacked = stacked
}
func (this *IType) SetCreator(creator Creator) {
	this.creator = creator
}

func (this *IType) SetListener(l Listener) {
	this.listener = l
}
