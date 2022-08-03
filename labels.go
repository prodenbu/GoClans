package GoClans

import "encoding/json"

/*
List player labels

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetClanLabels() ([]Label, error) {
	labels := LabelList{}
	body, err := s.get("labels/clans")
	if err != nil {
		return labels.Items, err
	}
	err = json.Unmarshal(body, &labels)
	if err != nil {
		return labels.Items, err
	}
	return labels.Items, nil
}

/*
List clan labels

optional limit int: Limit the number of items returned in the response.

optional after string: Return only items that occur after this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.

optional before string: Return only items that occur before this marker. Before marker can be found from the response,
inside the 'paging' property. Note that only after or before can be specified for a request, not both.
*/
func (s *Session) GetPlayerLabels() ([]Label, error) {
	labels := LabelList{}
	body, err := s.get("labels/players")
	if err != nil {
		return labels.Items, err
	}
	err = json.Unmarshal(body, &labels)
	if err != nil {
		return labels.Items, err
	}
	return labels.Items, nil
}
