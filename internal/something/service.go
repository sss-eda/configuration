package something

import "fmt"

// Service TODO
type Service struct{}

// AddSite TODO
func (service *Service) AddSite(
	site Site,
) error {
	return nil
}

// AddInstrument TODO
func (service *Service) AddInstrument(
	siteID SiteID,
	instrumentID InstrumentID,
) error {
	return nil
}

// ReconfigureInstrument TODO
func (service *Service) ReconfigureInstrument(
	instrumentID InstrumentID,
	newConfiguration Configuration,
) error {
	if !newConfiguration.Valid() {
		return fmt.Errorf("invalid configuration")
	}

	instrument := service.store.LoadInstrument(instrumentID)
	instrument.Reconfigure(newConfiguration)
}
