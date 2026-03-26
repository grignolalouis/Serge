package service

import (
	"github.com/serge-music/serge/internal/domain"
	"github.com/serge-music/serge/internal/repository"
)

type SupplierStats struct {
	domain.Supplier
	TotalOrders    int     `json:"total_orders"`
	ReceivedOrders int     `json:"received_orders"`
	OnTimeOrders   int     `json:"on_time_orders"`
	OnTimeRate     float64 `json:"on_time_rate"`
	AvgLeadDays    float64 `json:"avg_lead_days"`
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

func (s *ProcurementService) GetSupplierProduct(id string) (domain.SupplierProduct, error) {
	return s.repo.GetSupplierProduct(id)
}

func (s *ProcurementService) ListSupplierProducts() []domain.SupplierProduct {
	return s.repo.ListSupplierProducts()
}

func (s *ProcurementService) ListSupplierProductsBySupplier(supplierID string) []domain.SupplierProduct {
	return s.repo.ListSupplierProductsBySupplier(supplierID)
}

func (s *ProcurementService) ListSupplierProductsByProduct(productID string) []domain.SupplierProduct {
	return s.repo.ListSupplierProductsByProduct(productID)
}

func (s *ProcurementService) GetSupplierOrder(id string) (domain.SupplierOrder, error) {
	return s.repo.GetSupplierOrder(id)
}

func (s *ProcurementService) ListAllSupplierOrders() []domain.SupplierOrder {
	return s.repo.ListSupplierOrders()
}

func (s *ProcurementService) ListSupplierOrdersBySupplier(supplierID string) []domain.SupplierOrder {
	return s.repo.ListSupplierOrdersBySupplier(supplierID)
}

func (s *ProcurementService) ListSupplierOrdersByStatus(status domain.SupplierOrderStatus) []domain.SupplierOrder {
	return s.repo.ListSupplierOrdersByStatus(status)
}

func (s *ProcurementService) ListOverdueSupplierOrders() []domain.SupplierOrder {
	return s.repo.ListOverdueSupplierOrders()
}

func (s *ProcurementService) CompareSuppliers() []SupplierStats {
	suppliers := s.repo.ListSuppliers()
	var stats []SupplierStats
	for _, sup := range suppliers {
		orders := s.repo.ListSupplierOrdersBySupplier(sup.ID)
		st := SupplierStats{Supplier: sup, TotalOrders: len(orders)}
		var totalLeadDays int
		for _, so := range orders {
			if so.Status == domain.SupplierOrderReceived && so.ActualDelivery != nil {
				st.ReceivedOrders++
				lead := int(so.ActualDelivery.Sub(so.OrderDate).Hours() / 24)
				totalLeadDays += lead
				if !so.ActualDelivery.After(so.ExpectedDelivery) {
					st.OnTimeOrders++
				}
			}
		}
		if st.ReceivedOrders > 0 {
			st.OnTimeRate = float64(st.OnTimeOrders) / float64(st.ReceivedOrders)
			st.AvgLeadDays = float64(totalLeadDays) / float64(st.ReceivedOrders)
		}
		stats = append(stats, st)
	}
	return stats
}
