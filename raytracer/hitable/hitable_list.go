package hitable

import "CSC_4_sem_gowasm/raytracer"

type HitableList struct {
	List []Hitable
}

func NewHitableListSize(size uint32) *HitableList {
	return &HitableList{make([]Hitable, size)}
}

func NewHitableList(figures []Hitable) *HitableList {
	return &HitableList{figures}
}

func (figuresList *HitableList) Hit(ray raytracer.Ray, tMin, tMax float64) *raytracer.HitRecord {
	var hitRecord *raytracer.HitRecord = nil
	closest := tMax

	for _, figure := range figuresList.List {
		if hit := figure.Hit(ray, tMin, closest); hit != nil {
			closest = hit.T
			hitRecord = hit
		}
	}

	return hitRecord
}