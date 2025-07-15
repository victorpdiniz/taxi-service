# Testes de funcionalidade para o backend de corridas (PowerShell)

Write-Host "Teste: Corrida atrasada (monitorar)"
Invoke-WebRequest -Uri "http://localhost:3000/corrida/monitorar" -Method POST -Headers @{ "Content-Type" = "application/json" } -Body '{"MotoristaID":1,"PassageiroID":2,"TempoEstimado":20,"TempoDecorrido":25,"Preco":100.0,"Status":"em_andamento"}' | Select-Object -ExpandProperty Content

Write-Host "\nTeste: Corrida concluída com antecedência (finalizar)"
Invoke-WebRequest -Uri "http://localhost:3000/corrida/finalizar" -Method POST -Headers @{ "Content-Type" = "application/json" } -Body '{"MotoristaID":1,"PassageiroID":2,"TempoEstimado":20,"TempoDecorrido":15,"Preco":200.0,"Status":"em_andamento"}' | Select-Object -ExpandProperty Content

Write-Host "\nTeste: Corrida concluída no tempo previsto (finalizar)"
Invoke-WebRequest -Uri "http://localhost:3000/corrida/finalizar" -Method POST -Headers @{ "Content-Type" = "application/json" } -Body '{"MotoristaID":1,"PassageiroID":2,"TempoEstimado":20,"TempoDecorrido":20,"Preco":150.0,"Status":"em_andamento"}' | Select-Object -ExpandProperty Content

Write-Host "\nTeste: Corrida cancelada por excesso de tempo (cancelar)"
Invoke-WebRequest -Uri "http://localhost:3000/corrida/cancelar-por-excesso-tempo" -Method POST -Headers @{ "Content-Type" = "application/json" } -Body '{"MotoristaID":1,"PassageiroID":2,"TempoEstimado":20,"TempoDecorrido":40,"Preco":180.0,"Status":"em_andamento"}' | Select-Object -ExpandProperty Content
