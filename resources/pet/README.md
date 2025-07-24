
### Deletes a pet. <a name="delete"></a>

Delete a pet.

**API Endpoint**: `DELETE /pet/{petId}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `petId` | ✓ | Pet id to delete | `123` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	pet "pets_go/resources/pet"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Pet.Delete(pet.DeleteRequest{
		PetId: 123,
	})
}

```

### Finds Pets by status. <a name="find_by_status"></a>

Multiple status values can be provided with comma separated strings.

**API Endpoint**: `GET /pet/findByStatus`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `status` | ✗ | Status values that need to be considered for filter | `PetFindByStatusStatusEnumAvailable` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	pet "pets_go/resources/pet"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Pet.FindByStatus(pet.FindByStatusRequest{})
}

```

### Find pet by ID. <a name="get"></a>

Returns a single pet.

**API Endpoint**: `GET /pet/{petId}`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `petId` | ✓ | ID of pet to return | `123` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	pet "pets_go/resources/pet"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Pet.Get(pet.GetRequest{
		PetId: 123,
	})
}

```

### Add a new pet to the store. <a name="create"></a>

Add a new pet to the store.

**API Endpoint**: `POST /pet`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `name` | ✓ |  | `"doggie"` |
| `photoUrls` | ✓ |  | `[]string{"string",}` |
| `category` | ✗ |  | `Category {Id: nullable.NewValue(1),Name: nullable.NewValue("Dogs"),}` |
| `id` | ✗ |  | `10` |
| `status` | ✗ | pet status in the store | `PetStatusEnumAvailable` |
| `tags` | ✗ |  | `[]Tag{Tag {},}` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	nullable "pets_go/nullable"
	pet "pets_go/resources/pet"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Pet.Create(pet.CreateRequest{
		Id:   nullable.NewValue(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
	})
}

```

### Uploads an image. <a name="upload_image"></a>

Upload image of the pet.

**API Endpoint**: `POST /pet/{petId}/uploadImage`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `data` | ✓ |  | `sdkcore.MustOpenFile("uploads/file.pdf")` |
| `petId` | ✓ | ID of pet to update | `123` |
| `additionalMetadata` | ✗ | Additional Metadata | `"string"` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	sdkcore "pets_go/core"
	pet "pets_go/resources/pet"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Pet.UploadImage(pet.UploadImageRequest{
		Data:  sdkcore.MustOpenFile("uploads/file.pdf"),
		PetId: 123,
	})
}

```

#### Response

##### Type
[ApiResponse](/types/api_response.go)

##### Example
`ApiResponse {}`

### Update an existing pet. <a name="update"></a>

Update an existing pet by Id.

**API Endpoint**: `PUT /pet`

#### Parameters

| Parameter | Required | Description | Example |
|-----------|:--------:|-------------|--------|
| `name` | ✓ |  | `"doggie"` |
| `photoUrls` | ✓ |  | `[]string{"string",}` |
| `category` | ✗ |  | `Category {Id: nullable.NewValue(1),Name: nullable.NewValue("Dogs"),}` |
| `id` | ✗ |  | `10` |
| `status` | ✗ | pet status in the store | `PetStatusEnumAvailable` |
| `tags` | ✗ |  | `[]Tag{Tag {},}` |

#### Example Snippet

```go
package main

import (
	os "os"
	sdk "pets_go/client"
	nullable "pets_go/nullable"
	pet "pets_go/resources/pet"
)

func main() {
	client := sdk.NewClient(
		sdk.WithApiKey(os.Getenv("API_KEY")),
	)
	res, err := client.Pet.Update(pet.UpdateRequest{
		Id:   nullable.NewValue(10),
		Name: "doggie",
		PhotoUrls: []string{
			"string",
		},
	})
}

```
