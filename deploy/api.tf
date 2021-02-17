resource "aws_api_gateway_rest_api" "hello_universe" {
  name = "hello-universe"
}

resource "aws_api_gateway_deployment" "hello_universe" {
  triggers = {
    bff = md5(file("lambda.bff.tf"))
    bff = md5(file("lambda.tf"))
    bff = md5(file("api.tf"))
  }
  rest_api_id = aws_api_gateway_rest_api.hello_universe.id
  stage_name = "v${var.api_version}"
  stage_description = md5(file("api.tf"))

  lifecycle {
    create_before_destroy = true
  }
}


resource "aws_api_gateway_method" "root" {
  rest_api_id = aws_api_gateway_rest_api.hello_universe.id
  resource_id = aws_api_gateway_rest_api.hello_universe.root_resource_id
  http_method = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "root" {
  rest_api_id = aws_api_gateway_rest_api.hello_universe.id
  resource_id = aws_api_gateway_rest_api.hello_universe.root_resource_id
  http_method = aws_api_gateway_method.root.http_method
  type = "MOCK"

  request_parameters = {
    "integration.request.header.X-Authorization" = "'static'"
  }

  request_templates = {
    "application/xml" = <<EOF
{
   "body" : $input.json('$')
}
EOF
  }
}

