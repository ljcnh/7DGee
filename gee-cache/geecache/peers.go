/**
 * @Author: lj
 * @Description:
 * @File:  peers
 * @Version: 1.0.0
 * @Date: 2022/03/21 14:38
 */

package geecache

import pb "gee-cache/geecache/geecachepb"

type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(in *pb.Request, key *pb.Response) error
}
