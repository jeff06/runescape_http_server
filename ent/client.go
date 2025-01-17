// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"runescape_http_server/ent/migrate"

	"runescape_http_server/ent/skill"
	"runescape_http_server/ent/unlock"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Skill is the client for interacting with the Skill builders.
	Skill *SkillClient
	// Unlock is the client for interacting with the Unlock builders.
	Unlock *UnlockClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Skill = NewSkillClient(c.config)
	c.Unlock = NewUnlockClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Skill:  NewSkillClient(cfg),
		Unlock: NewUnlockClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Skill:  NewSkillClient(cfg),
		Unlock: NewUnlockClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Skill.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Skill.Use(hooks...)
	c.Unlock.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Skill.Intercept(interceptors...)
	c.Unlock.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *SkillMutation:
		return c.Skill.mutate(ctx, m)
	case *UnlockMutation:
		return c.Unlock.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// SkillClient is a client for the Skill schema.
type SkillClient struct {
	config
}

// NewSkillClient returns a client for the Skill from the given config.
func NewSkillClient(c config) *SkillClient {
	return &SkillClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `skill.Hooks(f(g(h())))`.
func (c *SkillClient) Use(hooks ...Hook) {
	c.hooks.Skill = append(c.hooks.Skill, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `skill.Intercept(f(g(h())))`.
func (c *SkillClient) Intercept(interceptors ...Interceptor) {
	c.inters.Skill = append(c.inters.Skill, interceptors...)
}

// Create returns a builder for creating a Skill entity.
func (c *SkillClient) Create() *SkillCreate {
	mutation := newSkillMutation(c.config, OpCreate)
	return &SkillCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Skill entities.
func (c *SkillClient) CreateBulk(builders ...*SkillCreate) *SkillCreateBulk {
	return &SkillCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SkillClient) MapCreateBulk(slice any, setFunc func(*SkillCreate, int)) *SkillCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SkillCreateBulk{err: fmt.Errorf("calling to SkillClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SkillCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SkillCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Skill.
func (c *SkillClient) Update() *SkillUpdate {
	mutation := newSkillMutation(c.config, OpUpdate)
	return &SkillUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SkillClient) UpdateOne(s *Skill) *SkillUpdateOne {
	mutation := newSkillMutation(c.config, OpUpdateOne, withSkill(s))
	return &SkillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SkillClient) UpdateOneID(id int) *SkillUpdateOne {
	mutation := newSkillMutation(c.config, OpUpdateOne, withSkillID(id))
	return &SkillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Skill.
func (c *SkillClient) Delete() *SkillDelete {
	mutation := newSkillMutation(c.config, OpDelete)
	return &SkillDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SkillClient) DeleteOne(s *Skill) *SkillDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SkillClient) DeleteOneID(id int) *SkillDeleteOne {
	builder := c.Delete().Where(skill.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SkillDeleteOne{builder}
}

// Query returns a query builder for Skill.
func (c *SkillClient) Query() *SkillQuery {
	return &SkillQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSkill},
		inters: c.Interceptors(),
	}
}

// Get returns a Skill entity by its id.
func (c *SkillClient) Get(ctx context.Context, id int) (*Skill, error) {
	return c.Query().Where(skill.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SkillClient) GetX(ctx context.Context, id int) *Skill {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUnlocks queries the unlocks edge of a Skill.
func (c *SkillClient) QueryUnlocks(s *Skill) *UnlockQuery {
	query := (&UnlockClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(skill.Table, skill.FieldID, id),
			sqlgraph.To(unlock.Table, unlock.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, skill.UnlocksTable, skill.UnlocksColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SkillClient) Hooks() []Hook {
	return c.hooks.Skill
}

// Interceptors returns the client interceptors.
func (c *SkillClient) Interceptors() []Interceptor {
	return c.inters.Skill
}

func (c *SkillClient) mutate(ctx context.Context, m *SkillMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SkillCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SkillUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SkillUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SkillDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Skill mutation op: %q", m.Op())
	}
}

// UnlockClient is a client for the Unlock schema.
type UnlockClient struct {
	config
}

// NewUnlockClient returns a client for the Unlock from the given config.
func NewUnlockClient(c config) *UnlockClient {
	return &UnlockClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `unlock.Hooks(f(g(h())))`.
func (c *UnlockClient) Use(hooks ...Hook) {
	c.hooks.Unlock = append(c.hooks.Unlock, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `unlock.Intercept(f(g(h())))`.
func (c *UnlockClient) Intercept(interceptors ...Interceptor) {
	c.inters.Unlock = append(c.inters.Unlock, interceptors...)
}

// Create returns a builder for creating a Unlock entity.
func (c *UnlockClient) Create() *UnlockCreate {
	mutation := newUnlockMutation(c.config, OpCreate)
	return &UnlockCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Unlock entities.
func (c *UnlockClient) CreateBulk(builders ...*UnlockCreate) *UnlockCreateBulk {
	return &UnlockCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UnlockClient) MapCreateBulk(slice any, setFunc func(*UnlockCreate, int)) *UnlockCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UnlockCreateBulk{err: fmt.Errorf("calling to UnlockClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UnlockCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UnlockCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Unlock.
func (c *UnlockClient) Update() *UnlockUpdate {
	mutation := newUnlockMutation(c.config, OpUpdate)
	return &UnlockUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UnlockClient) UpdateOne(u *Unlock) *UnlockUpdateOne {
	mutation := newUnlockMutation(c.config, OpUpdateOne, withUnlock(u))
	return &UnlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UnlockClient) UpdateOneID(id int) *UnlockUpdateOne {
	mutation := newUnlockMutation(c.config, OpUpdateOne, withUnlockID(id))
	return &UnlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Unlock.
func (c *UnlockClient) Delete() *UnlockDelete {
	mutation := newUnlockMutation(c.config, OpDelete)
	return &UnlockDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UnlockClient) DeleteOne(u *Unlock) *UnlockDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UnlockClient) DeleteOneID(id int) *UnlockDeleteOne {
	builder := c.Delete().Where(unlock.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UnlockDeleteOne{builder}
}

// Query returns a query builder for Unlock.
func (c *UnlockClient) Query() *UnlockQuery {
	return &UnlockQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUnlock},
		inters: c.Interceptors(),
	}
}

// Get returns a Unlock entity by its id.
func (c *UnlockClient) Get(ctx context.Context, id int) (*Unlock, error) {
	return c.Query().Where(unlock.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UnlockClient) GetX(ctx context.Context, id int) *Unlock {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUnlockIDSkillFk queries the unlock_id_skill_fk edge of a Unlock.
func (c *UnlockClient) QueryUnlockIDSkillFk(u *Unlock) *SkillQuery {
	query := (&SkillClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(unlock.Table, unlock.FieldID, id),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, unlock.UnlockIDSkillFkTable, unlock.UnlockIDSkillFkColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UnlockClient) Hooks() []Hook {
	return c.hooks.Unlock
}

// Interceptors returns the client interceptors.
func (c *UnlockClient) Interceptors() []Interceptor {
	return c.inters.Unlock
}

func (c *UnlockClient) mutate(ctx context.Context, m *UnlockMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UnlockCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UnlockUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UnlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UnlockDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Unlock mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Skill, Unlock []ent.Hook
	}
	inters struct {
		Skill, Unlock []ent.Interceptor
	}
)
