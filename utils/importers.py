import importlib
import pkgutil


def cargar_modulos(paquete):
    modulos = {}
    paquete_obj = importlib.import_module(paquete)

    for finder, nombre, es_paquete in pkgutil.iter_modules(paquete_obj.__path__):
        nombre_completo = f"{paquete}.{nombre}"
        modulos[nombre] = importlib.import_module(nombre_completo)

    return modulos

modulos = cargar_modulos("aoc2015")
