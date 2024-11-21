import time
from flask import Flask
from flask import request
import json # OH TERRYDAVIS PLEASE FORGIVE ME
app = Flask(__name__)

import random

def open_DB():
    x = open("db.json","r").read()
    x = json.loads(x)
    return x

def save_DB(x):
    x = json.dumps(x)
    open("db.json","w").write(x)

def generateRandom(N: int):
    o = []
    for _ in range(N):
        p = random.randint(ord("a"),ord("z"))
        o.append(chr(p))
    return "".join(o)

def GetUserByToken(token:str):
    x = open("db.json","r").read()
    x = json.loads(x)
    for i,u in enumerate(x["users"]): # may terry davis have mercy of this
        # time.sleep(random.random() / 2) # solution for side channel attack
        if u["token"] != token:
            continue
        # token = generateRandom(10)
        # x["users"][i]["token"] = token
        # return {"token":token, "userType":x["users"][i]["userType"]}
        return u
    return None

@app.post('/userSignIn')
def UserSignin():# this is so bad that if multiple requests to sigin the DB will just get fucked
    req = request.json
    req["user"]
    req["password"]
    x = open("db.json","r").read()
    x = json.loads(x)
    token = generateRandom(10)
    id_user = generateRandom(10)
    x["users"].append({"id": id_user, "user":req["user"], "password":req["password"], "token":token, "userType":0})
    open("db.json","w").write(json.dumps(x))
    return {"token":token, "userType":0}

@app.post('/Login')
def Login():
    req = request.json
    req["user"]
    req["password"]
    x = open("db.json","r").read()
    x = json.loads(x)
    for i,u in enumerate(x["users"]): # may terry davis have mercy of this
        # time.sleep(random.random() / 2) # solution for side channel attack
        if u["user"] != req["user"]:
            continue
        if u["password"] != req["password"]:
            continue
        token = generateRandom(10)
        x["users"][i]["token"] = token
        return {"token":token, "userType":x["users"][i]["userType"]}


@app.post('/CreateCampaigns')
def create_campaigns():# I'm tired
    req = request.json
    req["userToken"]
    req["name"]
    req["description"]
    user = GetUserByToken(req["userToken"])
    if not user:
        return
    if user["userType"] != 1: #architect
        return
    x = open("db.json","r").read()
    x = json.loads(x)
    x["campaigns"].append({"id": generateRandom(5),# why id is random? lie: non sequetial id is more secure. reality: function is already there, why not? i'm tire
                           "userToken":req["userToken"], "name": req["name"], "description":req["description"], "subscriptions":[]})
    open("db.json","w").write(json.dumps(x))


@app.post('/CreateEvents')
def CreateEvents():
    req = request.json
    campaign_id = req["CampaignId"]
    user_token = req["userToken"]
    date = req["date"]
    name = req["name"]
    local = req["local"]
    description = req["description"]
    user = GetUserByToken(req["userToken"])
    if not user:
        return
    if user["userType"] != 1: #architect
        return
    # open("campaign","+a").write({""})
    x = open("db.json","r").read()
    x = json.loads(x)
    if not any(True for c in x["campaigns"] if c["id"] == campaign_id):
        return
    token = generateRandom(10)
    x["Events"].append({"id": token,
                        "CampaignId":req["CampaignId"],
                       "userToken":req["userToken"],
                       "date": req["date"],
                       "name": req["name"],
                       "local": req["local"],
                       "description": req["description"],
                       "subscriptions": []})
    open("db.json","w").write(json.dumps(x))

@app.post('/SubscribeCampaign')
def SubscribeCampaign():
    req = request.json
    userToken = req["userToken"]
    campaing = req["campaignId"]
    user = GetUserByToken(userToken)
    if not user:
        return
    if user["userType"] != 0: #architect
        return
    db = open_DB()
    for i, c in enumerate(db["campaigns"]):
        if c["id"] != campaing:
            continue
        if user["id"] in db["campaigns"][i]["subscriptions"]:
            return
        db["campaigns"][i]["subscriptions"].append(user["id"])
    save_DB(db)

@app.post('/UnSubscribeCampaign')
def UnSubscribeCampaign():
    req = request.json
    userToken = req["userToken"]
    campaing = req["campaignId"]
    user = GetUserByToken(userToken)
    if not user:
        return
    if user["userType"] != 0: #architect
        return
    db = open_DB()
    for i, c in enumerate(db["campaigns"]):
        if c["id"] != campaing:
            continue
        if user["id"] not in db["campaigns"][i]["subscriptions"]:
            return
        db["campaigns"][i]["subscriptions"].remove(user["id"])
    save_DB(db)

@app.post('/ApplyEvent')
def ApplyEvent():
    req = request.json
    userToken = req["userToken"]
    campaing = req["EventId"]
    user = GetUserByToken(userToken)
    if not user:
        return
    if user["userType"] != 0: #architect
        return
    db = open_DB()
    for i, c in enumerate(db["Events"]):
        if c["id"] != campaing:
            continue
        if user["id"] in db["Events"][i]["subscriptions"]:
            return
        db["Events"][i]["subscriptions"].append(user["id"])
    save_DB(db)

@app.post('/UnApplyEvent')
def UnApplyEvent():
    req = request.json
    userToken = req["userToken"]
    campaing = req["EventId"]
    user = GetUserByToken(userToken)
    if not user:
        return
    if user["userType"] != 0: #architect
        return
    db = open_DB()
    for i, c in enumerate(db["Events"]):
        if c["id"] != campaing:
            continue
        if user["id"] not in db["Events"][i]["subscriptions"]:
            return
        db["Events"][i]["subscriptions"].remove(user["id"])
    save_DB(db)