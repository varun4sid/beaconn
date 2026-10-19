package scraper

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func (s *Session) GetToken() (TokenData, error) {
	tokenResponse := TokenData{
		Status: "error",
	}
	tokenURL := s.baseURL + "/Hostel/QRCode/QRcodeGenerate"

	response, err := s.Client.Get(tokenURL)
	if err != nil {
		return tokenResponse, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return tokenResponse, fmt.Errorf("failed to get token, status code: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return tokenResponse, fmt.Errorf("failed to parse token page: %w", err)
	}

	token, exists := doc.Find("img[alt='QR Code']").Attr("src")
	if !exists {
		tokenResponse.Status = "failed"
		return tokenResponse, fmt.Errorf("failed to find token in the response")
	}

	var tokens []TokenDetails

	doc.Find("table.card-table tbody tr").Each(func(i int, row *goquery.Selection) {
		cells := row.Find("td")
		if cells.Length() >= 6 {
			tokenName := cells.Eq(2).Text()
			tokenQty := cells.Eq(5).Text()

			tokens = append(tokens, TokenDetails{
				Name:     tokenName,
				Quantity: tokenQty,
			})
		}
	})

	tokenResponse = TokenData{
		Status:       "success",
		QRCode:       parseEncodedToken(token),
		TokenDetails: tokens,
	}

	return tokenResponse, nil
}

func parseEncodedToken(encodedToken string) string {
	// Assuming the token is in the format "data:image/png;base64,<token>"
	const prefix = "data:image/png;base64,"
	if len(encodedToken) > len(prefix) && encodedToken[:len(prefix)] == prefix {
		return encodedToken[len(prefix):]
	}
	return ""
}
