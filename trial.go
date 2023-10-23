import (
    "os"
    "github.com/aspose-words-cloud/aspose-words-cloud-go/dev/api/models"
)

config, _ := models.NewConfiguration("config.json")
wordsApi, ctx, _ := api.CreateWordsApi(config)

doc, _ := os.Open("Input.docx")
options := map[string]interface{}{}
request := &models.ConvertDocumentRequest{
    Document: doc,
    Format: ToStringPointer("html"),
    Optionals: options,
}
convert = wordsApi.ConvertDocument(ctx, request)