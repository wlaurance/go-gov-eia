package eia

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"
)

func getApiKey() string {
	key := os.Getenv("API_KEY")
	if key == "" {
		panic("Please set API_KEY in the shell environment.")
	}
	return key
}

func getClient() EIAClient {
	return Client(getApiKey(), &http.Client{})
}

func TestCategories(t *testing.T) {
	client := getClient()
	c, err := client.Categories()
	if err != nil {
		t.Errorf("Categories error %s", err.Error())
	} else {
		el := 1
		if len(c.ChildCategories) < el {
			t.Errorf("Categories had length %d, not %d", len(c.ChildCategories), el)
		}
	}
}

func TestCategoriesById(t *testing.T) {
	client := getClient()
	weeklyRetailGasPricesByAreaCats, err := client.CategoriesById("240691")
	if err != nil {
		t.Errorf("Categories error %s", err.Error())
	} else {
		el := 1
		if len(weeklyRetailGasPricesByAreaCats.ChildCategories) < el {
			t.Errorf("weeklyRetailGasPricesByAreaCats had length %d, not %d", len(weeklyRetailGasPricesByAreaCats.ChildCategories), el)
		}
	}
}

func TestPetroleumPaddsWeeklyRegular(t *testing.T) {
	client := getClient()
	weeklyRetailGasPricesByAreaCats, err := client.CategoriesById("240691")
	if err != nil {
		t.Errorf("Week petrol padds error %s", err.Error())
		t.FailNow()
	}
	for _, cat := range weeklyRetailGasPricesByAreaCats.ChildCategories {
		if strings.Contains(cat.Name, "PADD") {
			f, err := client.CategoriesById(fmt.Sprintf("%d", cat.CategoryId))
			if err != nil {
				t.Errorf("Some error %s", err.Error())
			}
			for _, ser := range f.ChildSeries {
				fmt.Println(ser.SeriesId)
			}
		}
	}
}
