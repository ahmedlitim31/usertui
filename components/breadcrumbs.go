package components

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Breadcrumb creates a navigation header for subpages
func Breadcrumb(tabName, pageName string) tview.Primitive {
	// Format the page title (convert snake_case to Title Case)
	pageTitle := formatPageTitle(pageName)
	// Create breadcrumb text
	breadcrumbText := fmt.Sprintf(" [%s] > %s ", tabName, pageTitle)

	// Create the breadcrumb view
	breadcrumb := tview.NewTextView().
		SetText(breadcrumbText).
		SetTextColor(tcell.ColorLawnGreen)

	return breadcrumb
}

// formatPageTitle converts "create_user" to "Create New User"
func formatPageTitle(pageName string) string {
	// Special cases for common page names
	specialCases := map[string]string{
		"create_user":    "Create New User",
		"edit_user":      "Edit User",
		"delete_user":    "Delete User",
		"view_users":     "View Users",
		"create_group":   "Create New Group",
		"edit_group":     "Edit Group",
		"delete_group":   "Delete Group",
		"view_groups":    "View Groups",
		"manage_members": "Manage Members",
		"main_menu":      "Main Menu",
	}

	if title, ok := specialCases[pageName]; ok {
		return title
	}

	// Fallback: convert snake_case to Title Case
	words := strings.Split(pageName, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + word[1:]
		}
	}
	return strings.Join(words, " ")
}

// WrapPageWithBreadcrumb adds a breadcrumb header to the page
func WrapPageWithBreadcrumb(page tview.Primitive, tabName, pageName string) tview.Primitive {
	// Create a flex container
	container := tview.NewFlex().SetDirection(tview.FlexRow)
	// Create a top padding spacer above the breadcrumb
	titleSpacer := tview.NewBox().SetBackgroundColor(tcell.ColorBlack)
	container.AddItem(titleSpacer, 1, 0, false)
	// Add breadcrumb
	breadcrumb := Breadcrumb(tabName, pageName)
	container.AddItem(breadcrumb, 1, 0, false)

	// Add the actual page content with some padding
	contentFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	contentFlex.SetBackgroundColor(tcell.ColorBlack)
	contentFlex.AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorBlack), 1, 0, false) // top padding
	contentFlex.AddItem(page, 0, 1, true)
	contentFlex.AddItem(tview.NewBox().SetBackgroundColor(tcell.ColorBlack), 1, 0, false) // bottom padding

	container.AddItem(contentFlex, 0, 1, true)

	return container
}
