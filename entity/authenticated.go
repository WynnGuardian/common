package entity

import "time"

type AuthenticatedItem struct {
	Id           string         `json:"id"`
	Item         string         `json:"item_name"`
	Weight       float64        `json:"weight"`
	Position     int            `json:"position"`
	OwnerMC      string         `json:"owner_mc_uuid"`
	OwnerDC      string         `json:"owner_dc_id"`
	Stats        map[string]int `json:"stats"`
	WynnItem     *WynnItem      `json:"wynn_item"`
	LastRanked   time.Time      `json:"last_weighted"`
	PublicOwner  bool           `json:"public_owner"`
	TrackingCode string         `json:"tracking_code"`
	Bytes        string         `json:"bytes"`
}
