resource "aws_lambda_function" "bff" {
  function_name = "bff"
  filename = "dist/bff.zip"
  source_code_hash = filebase64sha256("dist/bff.zip")
  handler = "bff"
  role = aws_iam_role.lambda.arn
  runtime = "go1.x"

  environment {
    variables = {
      "API_VERSION" = var.api_version
    }
  }
}

resource "aws_lambda_permission" "bff" {
  statement_id = "AllowAPIGatewayInvoke"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.bff.arn
  principal = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_resource" "bff" {
  rest_api_id = aws_api_gateway_rest_api.hello_universe.id
  parent_id = aws_api_gateway_rest_api.hello_universe.root_resource_id
  path_part = "bff"
}

resource "aws_api_gateway_method" "bff" {
  rest_api_id = aws_api_gateway_rest_api.hello_universe.id
  resource_id = aws_api_gateway_resource.bff.id
  http_method = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "bff" {
  rest_api_id = aws_api_gateway_rest_api.hello_universe.id
  resource_id = aws_api_gateway_resource.bff.id
  http_method = aws_api_gateway_method.bff.http_method
  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = aws_lambda_function.bff.invoke_arn
}

resource "aws_cloudwatch_log_group" "bff" {
  name = "/aws/lambda/${aws_lambda_function.bff.function_name}"
  retention_in_days = 14
}