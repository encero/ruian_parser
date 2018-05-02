// Code generated by entc, DO NOT EDIT.

package addressplace

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/encero/ruian_parser/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// OrientationNumber applies equality check predicate on the "orientation_number" field. It's identical to OrientationNumberEQ.
func OrientationNumber(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberLetter applies equality check predicate on the "orientation_number_letter" field. It's identical to OrientationNumberLetterEQ.
func OrientationNumberLetter(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrientationNumberLetter), v))
	})
}

// Zip applies equality check predicate on the "zip" field. It's identical to ZipEQ.
func Zip(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldZip), v))
	})
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumber), v))
	})
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumber), v))
	})
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int32) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumber), v...))
	})
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int32) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumber), v...))
	})
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumber), v))
	})
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumber), v))
	})
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumber), v))
	})
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumber), v))
	})
}

// OrientationNumberEQ applies the EQ predicate on the "orientation_number" field.
func OrientationNumberEQ(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberNEQ applies the NEQ predicate on the "orientation_number" field.
func OrientationNumberNEQ(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberIn applies the In predicate on the "orientation_number" field.
func OrientationNumberIn(vs ...int32) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrientationNumber), v...))
	})
}

// OrientationNumberNotIn applies the NotIn predicate on the "orientation_number" field.
func OrientationNumberNotIn(vs ...int32) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrientationNumber), v...))
	})
}

// OrientationNumberGT applies the GT predicate on the "orientation_number" field.
func OrientationNumberGT(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberGTE applies the GTE predicate on the "orientation_number" field.
func OrientationNumberGTE(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberLT applies the LT predicate on the "orientation_number" field.
func OrientationNumberLT(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberLTE applies the LTE predicate on the "orientation_number" field.
func OrientationNumberLTE(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrientationNumber), v))
	})
}

// OrientationNumberIsNil applies the IsNil predicate on the "orientation_number" field.
func OrientationNumberIsNil() predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrientationNumber)))
	})
}

// OrientationNumberNotNil applies the NotNil predicate on the "orientation_number" field.
func OrientationNumberNotNil() predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrientationNumber)))
	})
}

// OrientationNumberLetterEQ applies the EQ predicate on the "orientation_number_letter" field.
func OrientationNumberLetterEQ(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterNEQ applies the NEQ predicate on the "orientation_number_letter" field.
func OrientationNumberLetterNEQ(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterIn applies the In predicate on the "orientation_number_letter" field.
func OrientationNumberLetterIn(vs ...string) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrientationNumberLetter), v...))
	})
}

// OrientationNumberLetterNotIn applies the NotIn predicate on the "orientation_number_letter" field.
func OrientationNumberLetterNotIn(vs ...string) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrientationNumberLetter), v...))
	})
}

// OrientationNumberLetterGT applies the GT predicate on the "orientation_number_letter" field.
func OrientationNumberLetterGT(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterGTE applies the GTE predicate on the "orientation_number_letter" field.
func OrientationNumberLetterGTE(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterLT applies the LT predicate on the "orientation_number_letter" field.
func OrientationNumberLetterLT(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterLTE applies the LTE predicate on the "orientation_number_letter" field.
func OrientationNumberLetterLTE(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterContains applies the Contains predicate on the "orientation_number_letter" field.
func OrientationNumberLetterContains(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterHasPrefix applies the HasPrefix predicate on the "orientation_number_letter" field.
func OrientationNumberLetterHasPrefix(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterHasSuffix applies the HasSuffix predicate on the "orientation_number_letter" field.
func OrientationNumberLetterHasSuffix(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterIsNil applies the IsNil predicate on the "orientation_number_letter" field.
func OrientationNumberLetterIsNil() predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrientationNumberLetter)))
	})
}

// OrientationNumberLetterNotNil applies the NotNil predicate on the "orientation_number_letter" field.
func OrientationNumberLetterNotNil() predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrientationNumberLetter)))
	})
}

// OrientationNumberLetterEqualFold applies the EqualFold predicate on the "orientation_number_letter" field.
func OrientationNumberLetterEqualFold(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrientationNumberLetter), v))
	})
}

// OrientationNumberLetterContainsFold applies the ContainsFold predicate on the "orientation_number_letter" field.
func OrientationNumberLetterContainsFold(v string) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrientationNumberLetter), v))
	})
}

// ZipEQ applies the EQ predicate on the "zip" field.
func ZipEQ(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldZip), v))
	})
}

// ZipNEQ applies the NEQ predicate on the "zip" field.
func ZipNEQ(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldZip), v))
	})
}

// ZipIn applies the In predicate on the "zip" field.
func ZipIn(vs ...int32) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldZip), v...))
	})
}

// ZipNotIn applies the NotIn predicate on the "zip" field.
func ZipNotIn(vs ...int32) predicate.AddressPlace {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AddressPlace(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldZip), v...))
	})
}

// ZipGT applies the GT predicate on the "zip" field.
func ZipGT(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldZip), v))
	})
}

// ZipGTE applies the GTE predicate on the "zip" field.
func ZipGTE(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldZip), v))
	})
}

// ZipLT applies the LT predicate on the "zip" field.
func ZipLT(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldZip), v))
	})
}

// ZipLTE applies the LTE predicate on the "zip" field.
func ZipLTE(v int32) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldZip), v))
	})
}

// HasStreets applies the HasEdge predicate on the "streets" edge.
func HasStreets() predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StreetsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StreetsTable, StreetsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStreetsWith applies the HasEdge predicate on the "streets" edge with a given conditions (other predicates).
func HasStreetsWith(preds ...predicate.Street) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(StreetsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, StreetsTable, StreetsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AddressPlace) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AddressPlace) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AddressPlace) predicate.AddressPlace {
	return predicate.AddressPlace(func(s *sql.Selector) {
		p(s.Not())
	})
}
