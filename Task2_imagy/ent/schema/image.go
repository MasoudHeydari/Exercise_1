package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable(),
		field.String("original_url").NotEmpty(),
		field.String("local_name").NotEmpty().Unique(),
		field.String("file_extension").NotEmpty(),
		field.Int64("file_size"),
		field.Time("download_date").Default(time.Now),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
