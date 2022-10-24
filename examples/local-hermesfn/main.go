
package main

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func main() {
	config := openai.DefaultConfig("")
	config.BaseURL = "http://localhost:8080/oai/v1"

	model := "hermes-2-pro"

	client := openai.NewClientWithConfig(config)

	toolGetStockFundamentals := openai.Tool{
		Type: openai.ToolTypeFunction,

		Function: &openai.FunctionDefinition{
			Name: "get_stock_fundamentals",
			Description: `get_stock_fundamentals(symbol: str) -> dict - Get fundamental data for a given stock symbol using yfinance API.
    Args:
        symbol (str): The stock symbol.
    
    Returns:
        dict: A dictionary containing fundamental data.
            Keys:
                - 'symbol': The stock symbol.
                - 'company_name': The long name of the company.
                - 'sector': The sector to which the company belongs.
                - 'industry': The industry to which the company belongs.
                - 'market_cap': The market capitalization of the company.
                - 'pe_ratio': The forward price-to-earnings ratio.
                - 'pb_ratio': The price-to-book ratio.
                - 'dividend_yield': The dividend yield.