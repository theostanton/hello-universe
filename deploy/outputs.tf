//output "api_connection" {
//  value = aws_api_gateway_rest_api.hello_universe.
//}


// https://{restapi_id}.execute-api.{region}.amazonaws.com/{stage_name}/
output "api_base_url" {
  value = aws_api_gateway_deployment.hello_universe.invoke_url
}

output "api_bff_url" {
  value = "${aws_api_gateway_deployment.hello_universe.invoke_url}/${aws_api_gateway_resource.bff.path_part}"
}

