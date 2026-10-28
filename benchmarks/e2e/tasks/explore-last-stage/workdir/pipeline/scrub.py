from pipeline.config import SCRUB_BOOST
from pipeline.registry import register

_BASE = 26


def _scrub():
    return "orrick"


register("orrick", _BASE + SCRUB_BOOST, _scrub)
