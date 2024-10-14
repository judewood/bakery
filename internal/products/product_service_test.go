package products

import (
	"errors"
	"reflect"
	"testing"

	"github.com/judewood/bakery/myfmt"
)

var sampleProducts = []Product{
	{Name: "Vanilla cake", RecipeID: "1"},
	{Name: "plain cookie", RecipeID: "2"},
	{Name: "Doughnut", RecipeID: "3"},
}

func TestProductServiceGetAll(t *testing.T) {
	type testCase struct {
		name string
		want []Product
		err  error
	}

	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{name: "store error", want: nil, err: mockError},
		{name: "zero", want: []Product{}, err: nil},
		{name: "some", want: sampleProducts, err: nil},
	}

	t.Log("Given that I need to get all products")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("\n test %d: When store has %s products and has error is %v", i, test.name, test.err != nil)
				{
					mockProductStore := new(MockProductStore)
					mockProductStore.On("GetAll").Return(test.want, test.err)

					productService := NewProductService(mockProductStore)

					gotProducts, gotError := productService.GetAll()
					t.Logf("Then I get %v", test.want)
					{
						if reflect.DeepEqual(gotProducts, test.want) {
							t.Log(myfmt.ThumbsUp)
						} else {
							myfmt.Errorf(t, "\n%s Got %v", myfmt.ThumbsDown, gotProducts)
						}
					}
					t.Logf("Then I get error %v", test.err)
					{
						if test.err == nil {
							if gotError == nil {
								t.Log(myfmt.ThumbsUp)
							} else {
								t.Errorf("\n%s Got Error %v", myfmt.ThumbsDown, gotError.Error())
							}
						} else {
							if gotError == nil {
								t.Errorf("\n%s got no error ", myfmt.ThumbsDown)
							} else {
								t.Log(myfmt.ThumbsUp)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductServiceGet(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  Product
		err   error
	}

	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{name: "empty", input: "invalid", want: Product{}, err: mockError},
		{name: "valid", input: "Vanilla cake", want: sampleProducts[0], err: nil},
	}

	t.Log("Given that I need to get a product by its id")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("\n test %d: When I try to get a product that is %s", i, test.name)
				{
					mockProductStore := new(MockProductStore)
					mockProductStore.On("Get").Return(test.want, test.err)

					productService := NewProductService(mockProductStore)

					got, gotError := productService.Get(test.input)
					t.Logf("Then I get returned product: %#v", test.want)
					{
						if reflect.DeepEqual(got, test.want) {
							t.Log(myfmt.ThumbsUp)
						} else {
							t.Errorf("\n%s Got %v", myfmt.ThumbsDown, got)
						}
					}
					errStr := "an"
					if test.err == nil {
						errStr = "no"
					}
					t.Logf("And I get %s error", errStr)
					{
						if test.err == nil {
							if gotError == nil {
								t.Log(myfmt.ThumbsUp)
							} else {
								t.Errorf("\n%s got %v", myfmt.ThumbsDown, gotError)
							}
						} else {
							if gotError == nil {
								t.Errorf("\n%s got ", myfmt.ThumbsDown)
							} else {
								t.Log(myfmt.ThumbsUp)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductServiceAdd(t *testing.T) {
	type testCase struct {
		name  string
		input Product
		want  Product
		err   error
	}

	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{name: "empty", input: Product{}, want: Product{}, err: mockError},
		{name: "missing recipe id", input: Product{Name: "productName", RecipeID: ""}, want: Product{}, err: mockError},
		{name: "missing product name", input: Product{Name: "", RecipeID: "RecipeId"}, want: Product{}, err: mockError},
		{name: "valid", input: sampleProducts[0], want: sampleProducts[0], err: nil},
	}

	t.Log("Given that I need to add a product")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("\n test %d: When I try to add a product that is %s", i, test.name)
				{
					mockProductStore := new(MockProductStore)
					mockProductStore.On("Add").Return(test.want, test.err)

					productService := NewProductService(mockProductStore)

					got, gotError := productService.Add(test.input)
					t.Logf("Then I get returned product: %#v", test.want)
					{
						if reflect.DeepEqual(got, test.want) {
							t.Log(myfmt.ThumbsUp)
						} else {
							t.Errorf("\n%s Got %v", myfmt.ThumbsDown, got)
						}
					}
					errStr := "an"
					if test.err == nil {
						errStr = "no"
					}
					t.Logf("And I get %s error", errStr)
					{
						if test.err == nil {
							if gotError == nil {
								t.Log(myfmt.ThumbsUp)
							} else {
								t.Errorf("\n%s got %v", myfmt.ThumbsDown, gotError)
							}
						} else {
							if gotError == nil {
								t.Errorf("\n%s got ", myfmt.ThumbsDown)
							} else {
								t.Log(myfmt.ThumbsUp)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductServiceUpdate(t *testing.T) {
	type testCase struct {
		name  string
		input Product
		want  Product
		err   error
	}

	updatedProduct := sampleProducts[0]
	updatedProduct.Name = "updated name"
	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{name: "empty", input: Product{}, want: Product{}, err: mockError},
		{name: "missing recipe id", input: Product{Name: "productName", RecipeID: ""}, want: Product{}, err: mockError},
		{name: "missing product name", input: Product{Name: "", RecipeID: "RecipeId"}, want: Product{}, err: mockError},
		{name: "valid", input: updatedProduct, want: updatedProduct, err: nil},
	}

	t.Log("Given that I need to update a product")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("\n test %d: When I try to update a product that is %s", i, test.name)
				{
					mockProductStore := new(MockProductStore)
					mockProductStore.On("Add").Return(test.input, test.err)

					productService := NewProductService(mockProductStore)

					got, gotError := productService.Add(test.input)
					t.Logf("Then I get returned product: %#v", test.want)
					{
						if reflect.DeepEqual(got, test.want) {
							t.Log(myfmt.ThumbsUp)
						} else {
							t.Errorf("\n%s Got %v", myfmt.ThumbsDown, got)
						}
					}
					errStr := "an"
					if test.err == nil {
						errStr = "no"
					}
					t.Logf("And I get %s error", errStr)
					{
						if test.err == nil {
							if gotError == nil {
								t.Log(myfmt.ThumbsUp)
							} else {
								t.Errorf("\n%s got %v", myfmt.ThumbsDown, gotError)
							}
						} else {
							if gotError == nil {
								t.Errorf("\n%s got ", myfmt.ThumbsDown)
							} else {
								t.Log(myfmt.ThumbsUp)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestProductServiceDelete(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  Product
		err   error
	}

	mockError := errors.New("Mocked error")
	testCases := []testCase{
		{name: "does not exist", input: "invalid", want: Product{}, err: mockError},
		{name: "exists", input: "Vanilla cake", want: sampleProducts[0], err: nil},
	}

	t.Log("Given that I need to delete a product")
	{
		for i, test := range testCases {
			tf := func(t *testing.T) {
				t.Logf("\n test %d: When I try to delete a product that %s", i, test.name)
				{
					mockProductStore := new(MockProductStore)
					mockProductStore.On("Delete").Return(test.want, test.err)

					productService := NewProductService(mockProductStore)

					got, gotError := productService.Delete(test.input)
					t.Logf("Then I get returned product: %#v", test.want)
					{
						if reflect.DeepEqual(got, test.want) {
							t.Log(myfmt.ThumbsUp)
						} else {
							t.Errorf("\n%s Got %v", myfmt.ThumbsDown, got)
						}
					}
					errStr := "an"
					if test.err == nil {
						errStr = "no"
					}
					t.Logf("And I get %s error", errStr)
					{
						if test.err == nil {
							if gotError == nil {
								t.Log(myfmt.ThumbsUp)
							} else {
								t.Errorf("\n%s got %v", myfmt.ThumbsDown, gotError)
							}
						} else {
							if gotError == nil {
								t.Errorf("\n%s got ", myfmt.ThumbsDown)
							} else {
								t.Log(myfmt.ThumbsUp)
							}
						}
					}
				}
			}
			t.Run(test.name, tf)
		}
	}
}

func TestFormatProducts(t *testing.T) {
	want := "We have available:\n Vanilla cake\n plain cookie\n Doughnut\n"

	got := FormatProducts(sampleProducts)

	if got != want {
		myfmt.Errorf(t, "Failed TestFormatProducts. \nWanted:\n *%v*\nGot:\n *%v*", want, got)
	}
}
