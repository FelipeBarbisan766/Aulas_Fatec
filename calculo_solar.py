def calcular_energia_mensal(potencia_watts, tempo_horas, rendimento):
    """
    Calcula a geração de energia mensal de um módulo solar.
    
    :param potencia_watts: Potência do módulo solar em watts (W)
    :param tempo_horas: Tempo de exposição ao sol por dia em horas (h)
    :param rendimento: Rendimento do sistema (fração decimal, ex: 0.8 para 80%)
    :return: Energia gerada em um mês (kWh)
    """
    energia_diaria = potencia_watts * tempo_horas * rendimento  # Energia diária em Wh
    energia_mensal = (energia_diaria * 30) / 1000  # Conversão para kWh
    return energia_mensal
def calcular_valor(energia_mensal):
    """
    Calcula o valor da energia gerada mensalmente.
    
    :param energia_mensal: Energia gerada em um mês (kWh)
    :return: Valor da energia gerada mensalmente (R$)
    """
    preco_kwh = 0.671  # Preço do kWh em R$
    valor = energia_mensal * preco_kwh
    return valor
# Exemplo de uso
taxa_energia = calcular_energia_mensal(500, 12, 0.85)
valor = calcular_valor(calcular_energia_mensal(500, 12, 0.85))
print(f"Energia gerada no mês: {taxa_energia:.2f} kWh")
print(f"Valor da energia gerada no mês: R$ {valor:.2f}")