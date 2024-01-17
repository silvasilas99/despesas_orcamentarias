package model

import (
	"fmt"
)

type ExpenseResponse []struct {
	AnoMovimentacao             string `json: "ano_movimentacao"`
	MesMovimentacao             string `json: "mes_movimentacao"`
	orgao_codigo                string `json: "orgao_codigo"`
	orgao_nome                  string `json: "orgao_nome"`
	unidade_codigo              string `json: "unidade_codigo"`
	unidade_nome                string `json: "unidade_nome"`
	categoria_economica_codigo  string `json: "categoria_economica_codigo"`
	categoria_economica_nome    string `json: "categoria_economica_nome"`
	grupo_despesa_codigo        string `json: "grupo_despesa_codigo"`
	grupo_despesa_nome          string `json: "grupo_despesa_nome"`
	modalidade_aplicacao_codigo string `json: "modalidade_aplicacao_codigo"`
	modalidade_aplicacao_nome   string `json: "modalidade_aplicacao_nome"`
	elemento_codigo             string `json: "elemento_codigo"`
	elemento_nome               string `json: "elemento_nome"`
	subelemento_codigo          string `json: "subelemento_codigo"`
	subelemento_nome            string `json: "subelemento_nome"`
	funcao_codigo               string `json: "funcao_codigo"`
	funcao_nome                 string `json: "funcao_nome"`
	subfuncao_codigo            string `json: "subfuncao_codigo"`
	subfuncao_nome              string `json: "subfuncao_nome"`
	programa_codigo             string `json: "programa_codigo"`
	programa_nome               string `json: "programa_nome"`
	acao_codigo                 string `json: "acao_codigo"`
	acao_nome                   string `json: "acao_nome"`
	fonte_recurso_codigo        string `json: "fonte_recurso_codigo"`
	fonte_recurso_nome          string `json: "fonte_recurso_nome"`
	empenho_ano                 string `json: "empenho_ano"`
	empenho_modalidade_nome     string `json: "empenho_modalidade_nome"`
	empenho_modalidade_codigo   string `json: "empenho_modalidade_codigo"`
	empenho_numero              string `json: "empenho_numero"`
	subempenho                  string `json: "subempenho"`
	indicador_subempenho        string `json: "indicador_subempenho"`
	credor_codigo               string `json: "credor_codigo"`
	credor_nome                 string `json: "credor_nome"`
	modalidade_licitacao_codigo string `json: "modalidade_licitacao_codigo"`
	modalidade_licitacao_nome   string `json: "modalidade_licitacao_nome"`
	valor_empenhado             string `json: "valor_empenhado"`
	valor_liquidado             string `json: "valor_liquidado"`
	valor_pag                   string `json: "valor_pag"`
}

// TextOutput is exported,it formats the data to plain text.
func (c ExpenseResponse) TextOutput() string {
	p := fmt.Sprintf("AnoMovimentacao: %s\nMesMovimentacao : %s", c[0].AnoMovimentacao, c[0].MesMovimentacao)
	return p
}
