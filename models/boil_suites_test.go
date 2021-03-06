// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("ProductCategories", testProductCategories)
	t.Run("Products", testProducts)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesDelete)
	t.Run("Products", testProductsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesQueryDeleteAll)
	t.Run("Products", testProductsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesSliceDeleteAll)
	t.Run("Products", testProductsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesExists)
	t.Run("Products", testProductsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesFind)
	t.Run("Products", testProductsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesBind)
	t.Run("Products", testProductsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesOne)
	t.Run("Products", testProductsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesAll)
	t.Run("Products", testProductsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesCount)
	t.Run("Products", testProductsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesHooks)
	t.Run("Products", testProductsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesInsert)
	t.Run("ProductCategories", testProductCategoriesInsertWhitelist)
	t.Run("Products", testProductsInsert)
	t.Run("Products", testProductsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ProductToProductCategoryUsingCategory", testProductToOneProductCategoryUsingCategory)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ProductCategoryToCategoryProducts", testProductCategoryToManyCategoryProducts)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ProductToProductCategoryUsingCategoryProducts", testProductToOneSetOpProductCategoryUsingCategory)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ProductCategoryToCategoryProducts", testProductCategoryToManyAddOpCategoryProducts)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesReload)
	t.Run("Products", testProductsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesReloadAll)
	t.Run("Products", testProductsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesSelect)
	t.Run("Products", testProductsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesUpdate)
	t.Run("Products", testProductsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("ProductCategories", testProductCategoriesSliceUpdateAll)
	t.Run("Products", testProductsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
