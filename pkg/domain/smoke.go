package domain

import "github.com/graphql-go/graphql"

// Smoke represents an active smoke in the world.
type Smoke struct {
	Id       int      `json:"id"`
	Position Position `json:"position"`
}

func RenderSmoke(template map[string]interface{}, s Smoke) map[string]interface{} {
	result := map[string]interface{}{}

	_, hasId := template["id"]
	positionTemplate, hasPosition := template["position"]

	if hasId {
		result["id"] = s.Id
	}

	if hasPosition {
		tpl, _ := positionTemplate.(map[string]interface{})
		result["position"] = RenderPosition(tpl, s.Position)
	}

	return result
}

func RenderSmokes(template map[string]interface{}, s []Smoke) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	for _, v := range s {
		result = append(result, RenderSmoke(template, v))
	}

	return result
}

// SmokeType represents the GraphQL Type of a Smoke
var SmokeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "smoke",
	Fields: graphql.Fields{
		"position": &graphql.Field{
			Type: PositionType,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

// Remove removes a Smoke Id from smokes
func Remove(smokes []Smoke, toRemove int) []Smoke {
	result := make([]Smoke, 0)
	for _, single := range smokes {
		if single.Id != toRemove {
			result = append(result, single)
		}
	}

	return result
}
