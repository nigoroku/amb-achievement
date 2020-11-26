// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package generated

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTags)
	t.Run("InputAchievements", testInputAchievements)
	t.Run("MCategories", testMCategories)
	t.Run("OutputAchievementTags", testOutputAchievementTags)
	t.Run("OutputAchievements", testOutputAchievements)
	t.Run("TodoDetails", testTodoDetails)
	t.Run("Todos", testTodos)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsDelete)
	t.Run("InputAchievements", testInputAchievementsDelete)
	t.Run("MCategories", testMCategoriesDelete)
	t.Run("OutputAchievementTags", testOutputAchievementTagsDelete)
	t.Run("OutputAchievements", testOutputAchievementsDelete)
	t.Run("TodoDetails", testTodoDetailsDelete)
	t.Run("Todos", testTodosDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsQueryDeleteAll)
	t.Run("InputAchievements", testInputAchievementsQueryDeleteAll)
	t.Run("MCategories", testMCategoriesQueryDeleteAll)
	t.Run("OutputAchievementTags", testOutputAchievementTagsQueryDeleteAll)
	t.Run("OutputAchievements", testOutputAchievementsQueryDeleteAll)
	t.Run("TodoDetails", testTodoDetailsQueryDeleteAll)
	t.Run("Todos", testTodosQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsSliceDeleteAll)
	t.Run("InputAchievements", testInputAchievementsSliceDeleteAll)
	t.Run("MCategories", testMCategoriesSliceDeleteAll)
	t.Run("OutputAchievementTags", testOutputAchievementTagsSliceDeleteAll)
	t.Run("OutputAchievements", testOutputAchievementsSliceDeleteAll)
	t.Run("TodoDetails", testTodoDetailsSliceDeleteAll)
	t.Run("Todos", testTodosSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsExists)
	t.Run("InputAchievements", testInputAchievementsExists)
	t.Run("MCategories", testMCategoriesExists)
	t.Run("OutputAchievementTags", testOutputAchievementTagsExists)
	t.Run("OutputAchievements", testOutputAchievementsExists)
	t.Run("TodoDetails", testTodoDetailsExists)
	t.Run("Todos", testTodosExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsFind)
	t.Run("InputAchievements", testInputAchievementsFind)
	t.Run("MCategories", testMCategoriesFind)
	t.Run("OutputAchievementTags", testOutputAchievementTagsFind)
	t.Run("OutputAchievements", testOutputAchievementsFind)
	t.Run("TodoDetails", testTodoDetailsFind)
	t.Run("Todos", testTodosFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsBind)
	t.Run("InputAchievements", testInputAchievementsBind)
	t.Run("MCategories", testMCategoriesBind)
	t.Run("OutputAchievementTags", testOutputAchievementTagsBind)
	t.Run("OutputAchievements", testOutputAchievementsBind)
	t.Run("TodoDetails", testTodoDetailsBind)
	t.Run("Todos", testTodosBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsOne)
	t.Run("InputAchievements", testInputAchievementsOne)
	t.Run("MCategories", testMCategoriesOne)
	t.Run("OutputAchievementTags", testOutputAchievementTagsOne)
	t.Run("OutputAchievements", testOutputAchievementsOne)
	t.Run("TodoDetails", testTodoDetailsOne)
	t.Run("Todos", testTodosOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsAll)
	t.Run("InputAchievements", testInputAchievementsAll)
	t.Run("MCategories", testMCategoriesAll)
	t.Run("OutputAchievementTags", testOutputAchievementTagsAll)
	t.Run("OutputAchievements", testOutputAchievementsAll)
	t.Run("TodoDetails", testTodoDetailsAll)
	t.Run("Todos", testTodosAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsCount)
	t.Run("InputAchievements", testInputAchievementsCount)
	t.Run("MCategories", testMCategoriesCount)
	t.Run("OutputAchievementTags", testOutputAchievementTagsCount)
	t.Run("OutputAchievements", testOutputAchievementsCount)
	t.Run("TodoDetails", testTodoDetailsCount)
	t.Run("Todos", testTodosCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsHooks)
	t.Run("InputAchievements", testInputAchievementsHooks)
	t.Run("MCategories", testMCategoriesHooks)
	t.Run("OutputAchievementTags", testOutputAchievementTagsHooks)
	t.Run("OutputAchievements", testOutputAchievementsHooks)
	t.Run("TodoDetails", testTodoDetailsHooks)
	t.Run("Todos", testTodosHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsInsert)
	t.Run("InputAchievementTags", testInputAchievementTagsInsertWhitelist)
	t.Run("InputAchievements", testInputAchievementsInsert)
	t.Run("InputAchievements", testInputAchievementsInsertWhitelist)
	t.Run("MCategories", testMCategoriesInsert)
	t.Run("MCategories", testMCategoriesInsertWhitelist)
	t.Run("OutputAchievementTags", testOutputAchievementTagsInsert)
	t.Run("OutputAchievementTags", testOutputAchievementTagsInsertWhitelist)
	t.Run("OutputAchievements", testOutputAchievementsInsert)
	t.Run("OutputAchievements", testOutputAchievementsInsertWhitelist)
	t.Run("TodoDetails", testTodoDetailsInsert)
	t.Run("TodoDetails", testTodoDetailsInsertWhitelist)
	t.Run("Todos", testTodosInsert)
	t.Run("Todos", testTodosInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("InputAchievementTagToInputAchievementUsingInputAchievement", testInputAchievementTagToOneInputAchievementUsingInputAchievement)
	t.Run("InputAchievementTagToMCategoryUsingCategory", testInputAchievementTagToOneMCategoryUsingCategory)
	t.Run("InputAchievementToUserUsingUser", testInputAchievementToOneUserUsingUser)
	t.Run("OutputAchievementTagToOutputAchievementUsingOutputAchievement", testOutputAchievementTagToOneOutputAchievementUsingOutputAchievement)
	t.Run("OutputAchievementTagToMCategoryUsingCategory", testOutputAchievementTagToOneMCategoryUsingCategory)
	t.Run("OutputAchievementToUserUsingUser", testOutputAchievementToOneUserUsingUser)
	t.Run("TodoDetailToTodoUsingTodo", testTodoDetailToOneTodoUsingTodo)
	t.Run("TodoToUserUsingUser", testTodoToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("InputAchievementToInputAchievementTags", testInputAchievementToManyInputAchievementTags)
	t.Run("MCategoryToCategoryInputAchievementTags", testMCategoryToManyCategoryInputAchievementTags)
	t.Run("MCategoryToCategoryOutputAchievementTags", testMCategoryToManyCategoryOutputAchievementTags)
	t.Run("OutputAchievementToOutputAchievementTags", testOutputAchievementToManyOutputAchievementTags)
	t.Run("TodoToTodoDetails", testTodoToManyTodoDetails)
	t.Run("UserToInputAchievements", testUserToManyInputAchievements)
	t.Run("UserToOutputAchievements", testUserToManyOutputAchievements)
	t.Run("UserToTodos", testUserToManyTodos)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("InputAchievementTagToInputAchievementUsingInputAchievementTags", testInputAchievementTagToOneSetOpInputAchievementUsingInputAchievement)
	t.Run("InputAchievementTagToMCategoryUsingCategoryInputAchievementTags", testInputAchievementTagToOneSetOpMCategoryUsingCategory)
	t.Run("InputAchievementToUserUsingInputAchievements", testInputAchievementToOneSetOpUserUsingUser)
	t.Run("OutputAchievementTagToOutputAchievementUsingOutputAchievementTags", testOutputAchievementTagToOneSetOpOutputAchievementUsingOutputAchievement)
	t.Run("OutputAchievementTagToMCategoryUsingCategoryOutputAchievementTags", testOutputAchievementTagToOneSetOpMCategoryUsingCategory)
	t.Run("OutputAchievementToUserUsingOutputAchievements", testOutputAchievementToOneSetOpUserUsingUser)
	t.Run("TodoDetailToTodoUsingTodoDetails", testTodoDetailToOneSetOpTodoUsingTodo)
	t.Run("TodoToUserUsingTodos", testTodoToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("TodoDetailToTodoUsingTodoDetails", testTodoDetailToOneRemoveOpTodoUsingTodo)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("InputAchievementToInputAchievementTags", testInputAchievementToManyAddOpInputAchievementTags)
	t.Run("MCategoryToCategoryInputAchievementTags", testMCategoryToManyAddOpCategoryInputAchievementTags)
	t.Run("MCategoryToCategoryOutputAchievementTags", testMCategoryToManyAddOpCategoryOutputAchievementTags)
	t.Run("OutputAchievementToOutputAchievementTags", testOutputAchievementToManyAddOpOutputAchievementTags)
	t.Run("TodoToTodoDetails", testTodoToManyAddOpTodoDetails)
	t.Run("UserToInputAchievements", testUserToManyAddOpInputAchievements)
	t.Run("UserToOutputAchievements", testUserToManyAddOpOutputAchievements)
	t.Run("UserToTodos", testUserToManyAddOpTodos)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("TodoToTodoDetails", testTodoToManySetOpTodoDetails)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("TodoToTodoDetails", testTodoToManyRemoveOpTodoDetails)
}

func TestReload(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsReload)
	t.Run("InputAchievements", testInputAchievementsReload)
	t.Run("MCategories", testMCategoriesReload)
	t.Run("OutputAchievementTags", testOutputAchievementTagsReload)
	t.Run("OutputAchievements", testOutputAchievementsReload)
	t.Run("TodoDetails", testTodoDetailsReload)
	t.Run("Todos", testTodosReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsReloadAll)
	t.Run("InputAchievements", testInputAchievementsReloadAll)
	t.Run("MCategories", testMCategoriesReloadAll)
	t.Run("OutputAchievementTags", testOutputAchievementTagsReloadAll)
	t.Run("OutputAchievements", testOutputAchievementsReloadAll)
	t.Run("TodoDetails", testTodoDetailsReloadAll)
	t.Run("Todos", testTodosReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsSelect)
	t.Run("InputAchievements", testInputAchievementsSelect)
	t.Run("MCategories", testMCategoriesSelect)
	t.Run("OutputAchievementTags", testOutputAchievementTagsSelect)
	t.Run("OutputAchievements", testOutputAchievementsSelect)
	t.Run("TodoDetails", testTodoDetailsSelect)
	t.Run("Todos", testTodosSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsUpdate)
	t.Run("InputAchievements", testInputAchievementsUpdate)
	t.Run("MCategories", testMCategoriesUpdate)
	t.Run("OutputAchievementTags", testOutputAchievementTagsUpdate)
	t.Run("OutputAchievements", testOutputAchievementsUpdate)
	t.Run("TodoDetails", testTodoDetailsUpdate)
	t.Run("Todos", testTodosUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("InputAchievementTags", testInputAchievementTagsSliceUpdateAll)
	t.Run("InputAchievements", testInputAchievementsSliceUpdateAll)
	t.Run("MCategories", testMCategoriesSliceUpdateAll)
	t.Run("OutputAchievementTags", testOutputAchievementTagsSliceUpdateAll)
	t.Run("OutputAchievements", testOutputAchievementsSliceUpdateAll)
	t.Run("TodoDetails", testTodoDetailsSliceUpdateAll)
	t.Run("Todos", testTodosSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}