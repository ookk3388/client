// Auto-generated by avdl-compiler v1.3.27 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/favorite.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type FolderType int

const (
	FolderType_UNKNOWN FolderType = 0
	FolderType_PRIVATE FolderType = 1
	FolderType_PUBLIC  FolderType = 2
	FolderType_TEAM    FolderType = 3
)

func (o FolderType) DeepCopy() FolderType { return o }

var FolderTypeMap = map[string]FolderType{
	"UNKNOWN": 0,
	"PRIVATE": 1,
	"PUBLIC":  2,
	"TEAM":    3,
}

var FolderTypeRevMap = map[FolderType]string{
	0: "UNKNOWN",
	1: "PRIVATE",
	2: "PUBLIC",
	3: "TEAM",
}

func (e FolderType) String() string {
	if v, ok := FolderTypeRevMap[e]; ok {
		return v
	}
	return ""
}

// Folder represents a favorite top-level folder in kbfs.
// This type is likely to change significantly as all the various parts are
// connected and tested.
type Folder struct {
	Name            string     `codec:"name" json:"name"`
	Private         bool       `codec:"private" json:"private"`
	NotificationsOn bool       `codec:"notificationsOn" json:"notificationsOn"`
	Created         bool       `codec:"created" json:"created"`
	FolderType      FolderType `codec:"folderType" json:"folderType"`
	TeamID          *TeamID    `codec:"team_id,omitempty" json:"team_id,omitempty"`
}

func (o Folder) DeepCopy() Folder {
	return Folder{
		Name:            o.Name,
		Private:         o.Private,
		NotificationsOn: o.NotificationsOn,
		Created:         o.Created,
		FolderType:      o.FolderType.DeepCopy(),
		TeamID: (func(x *TeamID) *TeamID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TeamID),
	}
}

type FavoritesResult struct {
	FavoriteFolders []Folder `codec:"favoriteFolders" json:"favoriteFolders"`
	IgnoredFolders  []Folder `codec:"ignoredFolders" json:"ignoredFolders"`
	NewFolders      []Folder `codec:"newFolders" json:"newFolders"`
}

func (o FavoritesResult) DeepCopy() FavoritesResult {
	return FavoritesResult{
		FavoriteFolders: (func(x []Folder) []Folder {
			if x == nil {
				return nil
			}
			ret := make([]Folder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.FavoriteFolders),
		IgnoredFolders: (func(x []Folder) []Folder {
			if x == nil {
				return nil
			}
			ret := make([]Folder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.IgnoredFolders),
		NewFolders: (func(x []Folder) []Folder {
			if x == nil {
				return nil
			}
			ret := make([]Folder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.NewFolders),
	}
}

type FavoriteAddArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Folder    Folder `codec:"folder" json:"folder"`
}

type FavoriteIgnoreArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Folder    Folder `codec:"folder" json:"folder"`
}

type GetFavoritesArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type FavoriteInterface interface {
	// Adds a folder to a user's list of favorite folders.
	FavoriteAdd(context.Context, FavoriteAddArg) error
	// Removes a folder from a user's list of favorite folders.
	FavoriteIgnore(context.Context, FavoriteIgnoreArg) error
	// Returns all of a user's favorite folders.
	GetFavorites(context.Context, int) (FavoritesResult, error)
}

func FavoriteProtocol(i FavoriteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.favorite",
		Methods: map[string]rpc.ServeHandlerDescription{
			"favoriteAdd": {
				MakeArg: func() interface{} {
					var ret [1]FavoriteAddArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]FavoriteAddArg)
					if !ok {
						err = rpc.NewTypeError((*[1]FavoriteAddArg)(nil), args)
						return
					}
					err = i.FavoriteAdd(ctx, typedArgs[0])
					return
				},
			},
			"favoriteIgnore": {
				MakeArg: func() interface{} {
					var ret [1]FavoriteIgnoreArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]FavoriteIgnoreArg)
					if !ok {
						err = rpc.NewTypeError((*[1]FavoriteIgnoreArg)(nil), args)
						return
					}
					err = i.FavoriteIgnore(ctx, typedArgs[0])
					return
				},
			},
			"getFavorites": {
				MakeArg: func() interface{} {
					var ret [1]GetFavoritesArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]GetFavoritesArg)
					if !ok {
						err = rpc.NewTypeError((*[1]GetFavoritesArg)(nil), args)
						return
					}
					ret, err = i.GetFavorites(ctx, typedArgs[0].SessionID)
					return
				},
			},
		},
	}
}

type FavoriteClient struct {
	Cli rpc.GenericClient
}

// Adds a folder to a user's list of favorite folders.
func (c FavoriteClient) FavoriteAdd(ctx context.Context, __arg FavoriteAddArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.favorite.favoriteAdd", []interface{}{__arg}, nil)
	return
}

// Removes a folder from a user's list of favorite folders.
func (c FavoriteClient) FavoriteIgnore(ctx context.Context, __arg FavoriteIgnoreArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.favorite.favoriteIgnore", []interface{}{__arg}, nil)
	return
}

// Returns all of a user's favorite folders.
func (c FavoriteClient) GetFavorites(ctx context.Context, sessionID int) (res FavoritesResult, err error) {
	__arg := GetFavoritesArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.favorite.getFavorites", []interface{}{__arg}, &res)
	return
}
