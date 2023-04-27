# SETUP

import datetime
import csv
import json
import os
from collections import defaultdict

import pandas as pd

import requests
from associado import Associado

# from pandas import json_normalize

user, password = 'guionardo', 'A11977'

HOST = 'http://localhost:5000'  # 'https://mappa-proxy.fly.dev'

# with open("mappa/mappa.json") as f:
# user,password=json.load(f)
auth, user_id = '', 0


def erro(*msg):
    print('ERROR', *msg)
    exit()


def get(url: str, title):
    print('GET', title, end='')
    url = url.removeprefix('/')
    response = requests.get((f'{HOST}/{url}'), headers={'Authorization': auth})
    if response.status_code > 299:
        erro(response, response.json())
    print(' ', response.reason)
    return response.json()


def login(user, password):
    global auth, user_id
    url = (f'{HOST}/mappa/login')
    if os.path.exists('auth.json'):
        with open('auth.json') as f:
            auth_json = json.load(f)
            created = datetime.datetime.strptime(
                auth_json['created'][:19], '%Y-%m-%dT%H:%M:%S')
            print('Auth created', created)

            if (created+datetime.timedelta(seconds=auth_json['ttl'])) > datetime.datetime.now():
                auth, user_id = auth_json['id'], auth_json['userId']
                print('CACHED LOGIN OK')
                return

    response = requests.post(url, json=dict(username=user, password=password))
    if response.status_code != 200:
        erro(response, response.json())
    auth_json = response.json()
    with open('auth.json', 'w') as f:
        json.dump(auth_json, f)
    auth, user_id = auth_json['id'], auth_json['userId']
    print('Login OK', auth_json)


print('Login mappa: ', user)

login(user, password)

# Progressões

progressoes = get(
    '/mappa/progressoes/L', 'PROGRESSÕES')

Associado.set_progressoes(progressoes)

escotista = get(f'/mappa/escotista/{user_id}', 'ESCOTISTA')

print('Escotista', escotista['associado']['nome'])
print('Grupo', escotista['grupo']['nome'],
      f'{escotista["grupo"]["codigo"]}/{escotista["grupo"]["codigoRegiao"]}')

secoes = get(f'mappa/secoes/{user_id}', 'SEÇõES')

marcacoes_secao = {}

associados = {}
desenvolvimentos = defaultdict(int)
csv_file = open('stats_secao.csv', 'w')
csv_writer = csv.DictWriter(csv_file, [
                            'secao', 'subsecao', 'associado', 'promessa', 'pata tenra', 'saltador', 'rastreador', 'cacador',
                            'data_passagem',
                            'semanas_restantes', 'desenv_por_idade', 'desenv_por_progressao', 'pct_idade', 'pct_progressao'])
csv_writer.writeheader()
for secao in secoes:
    print(f'{secao["codigo"]} {secao["nome"]}')
    codigo_secao = secao['codigo']
    marcacoes = get(f'/mappa/marcacoes/{codigo_secao}', 'MARCAÇÕES')
    for subsecao in secao["subsecoes"]:
        print(f'  {subsecao["nome"]}')
        for associado in subsecao['associados']:
            associado = Associado(associado, marcacoes)
            associados[associado.codigo] = associado
            print(f'     {associado}')
            (_, dev_idade) = associado.get_desenvolvimento_por_idade()
            (_, dev_progressao) = associado.get_desenvolvimento_por_marcacoes()
            csv_writer.writerow({'secao': secao['nome'],
                                 'subsecao': subsecao['nome'],
                                 'associado': associado.nome,
                                 'promessa': associado.desenvolvimentos['PROMESSA'],
                                 'pata tenra': associado.desenvolvimentos['PATA_TENRA'],
                                 'saltador': associado.desenvolvimentos['SALTADOR'],
                                 'rastreador': associado.desenvolvimentos['RASTREADOR'],
                                 'cacador': associado.desenvolvimentos['CAÇADOR'],
                                 'data_passagem': associado.data_passagem.date(),
                                 'semanas_restantes': associado.semanas_restantes,
                                 'desenv_por_idade': dev_idade,
                                 'desenv_por_progressao': dev_progressao,
                                 'pct_idade': associado.get_pct_idade(),
                                 'pct_progressao': associado.get_pct_progressao()})

    marcacoes_secao[codigo_secao] = marcacoes

    print('Seção', secao['codigo'], secao['nome'],
          len(marcacoes_secao[codigo_secao]))

csv_file.close()
df = pd.read_csv('stats_secao.csv')

print(df)
df.to_excel('stats_secao.xlsx', engine='xlsxwriter')
