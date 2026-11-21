package iyzilink

import "encoding/json"

// API'nin metin veya sayı olarak gönderdiği fiyatları okur.
func (d *DetailResponseData) UnmarshalJSON(data []byte) error {
	type detail DetailResponseData
	wire := struct {
		*detail
		Price             json.Number   `json:"price"`
		PresetPriceValues []json.Number `json:"presetPriceValues"`
	}{detail: (*detail)(d)}
	if err := json.Unmarshal(data, &wire); err != nil {
		return err
	}
	if wire.Price != "" {
		price, err := wire.Price.Float64()
		if err != nil {
			return err
		}
		d.Price = price
	}
	if wire.PresetPriceValues != nil {
		prices := make([]float64, len(wire.PresetPriceValues))
		for i, value := range wire.PresetPriceValues {
			price, err := value.Float64()
			if err != nil {
				return err
			}
			prices[i] = price
		}
		d.PresetPriceValues = prices
	}
	return nil
}
