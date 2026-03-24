package service

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/repository"
)

type SupplierStats struct {
	domain.Supplier
	TotalPOs    int     `json:"total_pos"`
	ReceivedPOs int     `json:"received_pos"`
	OnTimePOs   int     `json:"on_time_pos"`
	OnTimeRate  float64 `json:"on_time_rate"`
	AvgLeadDays float64 `json:"avg_lead_days"`
}

type ProcurementService struct {
	repo *repository.ProcurementRepository
}

func NewProcurementService(repo *repository.ProcurementRepository) *ProcurementService {
	return &ProcurementService{repo: repo}
}

func (s *ProcurementService) GetSupplier(id string) (domain.Supplier, error) {
	return s.repo.GetSupplier(id)
}

func (s *ProcurementService) ListSuppliers() []domain.Supplier {
	return s.repo.ListSuppliers()
}

func (s *ProcurementService) GetPurchaseOrder(id string) (domain.PurchaseOrder, error) {
	return s.repo.GetPurchaseOrder(id)
}

func (s *ProcurementService) ListPurchaseOrdersBySupplier(supplierID string) []domain.PurchaseOrder {
	return s.repo.ListPurchaseOrdersBySupplier(supplierID)
}

func (s *ProcurementService) ListPurchaseOrdersByStatus(status domain.PurchaseOrderStatus) []domain.PurchaseOrder {
	return s.repo.ListPurchaseOrdersByStatus(status)
}

func (s *ProcurementService) ListOverduePurchaseOrders() []domain.PurchaseOrder {
	return s.repo.ListOverduePurchaseOrders()
}

func (s *ProcurementService) CompareSuppliers() []SupplierStats {
	suppliers := s.repo.ListSuppliers()
	var stats []SupplierStats
	for _, sup := range suppliers {
		pos := s.repo.ListPurchaseOrdersBySupplier(sup.ID)
		st := SupplierStats{Supplier: sup, TotalPOs: len(pos)}
		var totalLeadDays int
		for _, po := range pos {
			if po.Status == domain.POReceived && po.ActualDelivery != nil {
				st.ReceivedPOs++
				lead := int(po.ActualDelivery.Sub(po.OrderDate).Hours() / 24)
				totalLeadDays += lead
				if !po.ActualDelivery.After(po.ExpectedDelivery) {
					st.OnTimePOs++
				}
			}
		}
		if st.ReceivedPOs > 0 {
			st.OnTimeRate = float64(st.OnTimePOs) / float64(st.ReceivedPOs)
			st.AvgLeadDays = float64(totalLeadDays) / float64(st.ReceivedPOs)
		}
		stats = append(stats, st)
	}
	return stats
}
