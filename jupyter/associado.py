from collections import defaultdict
from datetime import datetime, timedelta
from typing import Tuple


class Associado:
    PROGRESSOES = dict()
    DESENVOLVIMENTO = defaultdict(int)

    def __init__(self, associado: dict, marcacoes: list):
        self.codigo = associado['codigo']
        self.nome = associado['nome']
        self.data_nasc = datetime.strptime(
            associado['dataNascimento'][:10], '%Y-%m-%d')
        self.data_passagem = self.data_nasc + timedelta(days=365.25*11)
        self.semanas_restantes = round(
            (self.data_passagem - datetime.now()).days/7)
        self.desenvolvimentos = defaultdict(int)
        self.progressoes = []
        self.progressao_total = 0
        marcas = {m['codigoAtividade']: m
                  for m in marcacoes
                  if m['codigoAssociado'] == self.codigo}
        for (cod_progressao, progressao) in self.PROGRESSOES.items():
            if marcas.get(cod_progressao):
                _, desenv = self.get_desenvolvimento(
                    progressao['codigoCaminho'], progressao['codigoDesenvolvimento'])
                self.desenvolvimentos[desenv] += 1
                self.progressao_total += 1

        self.progressao_pct = round(100*self.progressao_total /
                                    sum(self.DESENVOLVIMENTO.values()))
        self.desenvolvimento_por_idade = self.get_desenvolvimento_por_idade()

    @classmethod
    def set_progressoes(cls, progressoes: list):
        cls.DESENVOLVIMENTO.clear()
        for p in progressoes:
            _, desenv = cls.get_desenvolvimento(
                p['codigoCaminho'], p['codigoDesenvolvimento'])
            cls.DESENVOLVIMENTO[desenv] += 1

        cls.PROGRESSOES = {p['codigo']: p for p in progressoes}

    @classmethod
    def get_desenvolvimento(cls, cod_caminho, cod_desenvolvimento) -> Tuple[int, str]:
        match cod_caminho:
            case 1:
                return 0, 'PROMESSA'
            case 2:
                if cod_desenvolvimento in [19, 20, 21]:
                    return 1, 'PATA_TENRA'
                return 2, 'SALTADOR'
            case 3:
                if cod_desenvolvimento in [19, 20, 21]:
                    return 3, 'RASTREADOR'
                return 4, 'CAÇADOR'

        return -1, '__ DESCONHECIDO __'

    def get_desenvolvimento_por_idade(self) -> Tuple[int, str]:
        """
        6.5 - 7 : PROMESSA
        7 - 8 : PATA_TENRA
        8 - 9 : SALTADOR
        9 - 10 : RASTREADOR
        10 - 11: CAÇADOR
        """
        idade = (datetime.now() - self.data_nasc).days / 365.25
        if idade < 7:
            return 0, 'PROMESSA'
        if idade < 8:
            return 1, 'PATA_TENRA'
        if idade < 9:
            return 2, 'SALTADOR'
        return (3, 'RASTREADOR') if idade < 10 else (4, 'CAÇADOR')

    def get_desenvolvimento_por_marcacoes(self) -> Tuple[int, str]:
        return next(
            (
                (4 - index, desenvolvimento)
                for index, desenvolvimento in enumerate(
                    reversed(self.DESENVOLVIMENTO.keys())
                )
                if self.desenvolvimentos[desenvolvimento] > 0
            ),
            (0, 'PROMESSA'),
        )

    def status(self) -> str:
        desenvolvimentos = [
            f'{d} = {self.desenvolvimentos[d]:02d}/{self.DESENVOLVIMENTO[d]}'
            for d in self.DESENVOLVIMENTO]
        return ', '.join(desenvolvimentos)

    def __str__(self) -> str:
        ida, desenv_atual = self.get_desenvolvimento_por_marcacoes()
        idi, desenv_idade = self.get_desenvolvimento_por_idade()
        if desenv_atual == desenv_idade:
            desenv = desenv_atual
        else:
            sinal = '<' if ida < idi else '>'
            desenv = f'{desenv_atual} {sinal} {desenv_idade}'
        return f'{self.nome:40s} {self.status()} ({desenv}) [{self.semanas_restantes} sem -> {self.progressao_pct}%]'

    def get_pct_idade(self) -> float:
        d0 = self.data_nasc+timedelta(days=365.25*6.5)
        d = (self.data_passagem - d0).days
        return round((datetime.now()-d0).days / d, 3)

    def get_pct_progressao(self) -> float:
        p = sum(self.desenvolvimentos.values())
        t = sum(self.DESENVOLVIMENTO.values())
        return round(p/t, 3)
