package gocassa

type oneToManyT struct {
	*t
	fieldToIndexBy string
	idField        string
}

func (o *oneToManyT) Update(field, id interface{}, m map[string]interface{}) Op {
	return o.Where(Eq(o.fieldToIndexBy, field), Eq(o.idField, id)).Update(m)
}

func (o *oneToManyT) UpdateWithOptions(field, id interface{}, m map[string]interface{}, opts Options) Op {
	return o.Where(Eq(o.fieldToIndexBy, field), Eq(o.idField, id)).UpdateWithOptions(m, opts)
}

func (o *oneToManyT) Delete(field, id interface{}) Op {
	return o.Where(Eq(o.fieldToIndexBy, field), Eq(o.idField, id)).Delete()
}

func (o *oneToManyT) DeleteAll(field interface{}) Op {
	return o.Where(Eq(o.fieldToIndexBy, field)).Delete()
}

func (o *oneToManyT) Read(field, id, pointer interface{}) Op {
	return o.Where(Eq(o.fieldToIndexBy, field), Eq(o.idField, id)).Query().ReadOne(pointer)
}

func (o *oneToManyT) MultiRead(field interface{}, ids []interface{}, pointerToASlice interface{}) Op {
	return o.Where(Eq(o.fieldToIndexBy, field), In(o.idField, ids...)).Query().Read(pointerToASlice)
}

func (o *oneToManyT) List(field, startId interface{}, limit int, pointerToASlice interface{}) Op {
	rels := []Relation{Eq(o.fieldToIndexBy, field)}
	if startId != nil {
		rels = append(rels, GTE(o.idField, startId))
	}
	return o.Where(rels...).Query().Limit(limit).Read(pointerToASlice)
}
