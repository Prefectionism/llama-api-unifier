#!/usr/bin/env python
import os
import uvicorn

from typing import Any, List, Union
from fastapi import FastAPI
from langserve import add_routes

from langchain import hub
from langchain_openai import ChatOpenAI
from langchain.agents import AgentExecutor, create_react_agent
from langchain_core.messages import AIMessage, FunctionMessage, HumanMessage

from langchain.pydantic_v1 import BaseModel, Field
from langchain_community.tools import DuckDuckGoSearchResults

llm = ChatOpenAI(model_name=os.environ['MODEL_NAME'], temperature=0, streaming=True)

# https://smith.langchain.com/hub/hwchase17/react-chat
prompt = hub.pull("hwchase17/react-chat")

tools = [DuckDuckGoSearchResults(max_results=1)]
agent = create_react_agent(llm, to