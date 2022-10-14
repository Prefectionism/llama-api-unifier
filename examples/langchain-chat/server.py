#!/usr/bin/env python
import os
import uvicorn

from typing import Any, List, Union
from fastapi import FastAPI
from langserve import add_routes

from langchain import hub
from langchain_openai import ChatOpenAI
from langchain.agents import